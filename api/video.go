package api

import (
	"encoding/json"
	"fmt"
	"io"
	"m3u82mp4/api/middleware"
	"m3u82mp4/consts"
	"m3u82mp4/consts/errcode"
	"m3u82mp4/library"
	"m3u82mp4/model/db"
	systemser "m3u82mp4/model/service/system"
	"m3u82mp4/model/video"
	"m3u82mp4/utils"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type VideoApi struct {
	BaseApi
}

// InitVideoRouter 初始化路由
func InitVideoRouter(r *gin.RouterGroup) {
	api := &VideoApi{}
	api.RouterGroup = r.Group("/video", middleware.SetUserInfo)
	api.ApiByte("GET", "/query", api.Query)       // 列表
	api.Api("GET", "/getInfoV2", api.GetInfo)     // 详情
	api.Api("GET", "/playRecord", api.PlayRecord) // 播放记录
	api.Api("POST", "/star", api.Star)            // 收藏
	api.Api("GET", "/starList", api.StarList)     // 收藏列表
	api.ApiFile("GET", "/download", api.Download) // 下载
}

// Query 查询影视信息
func (v *VideoApi) Query(c *gin.Context) any {
	res := make([]byte, 0)
	key := c.Query("key")
	if len(key) <= 1 {
		return errcode.VIDEO_QUERY_KEY_LEN_ERR
	}
	// 将编码后的字符串中的 '%' 替换为 '%25'
	encodedStr := strings.Replace(url.QueryEscape(consts.VIDEO_URL_SOURCE), "%", "%25", -1)
	// 第二次 URL 编码
	doubleEncodedStr := url.QueryEscape(encodedStr)
	resUrl := fmt.Sprintf(systemser.GetURL(c), url.QueryEscape(key), doubleEncodedStr)
	v.Logger.Info("video-query-url", zap.String("url", resUrl), zap.Int64("userid", v.User.Id))
	resp, err := http.Get(resUrl)
	if err != nil {
		v.Logger.Error("video-query-url-获取地址错误", zap.Error(err))
		return res
	}
	defer resp.Body.Close()
	bytedata, err := io.ReadAll(resp.Body)
	if err != nil {
		v.Logger.Error("video-query-url-解析数据错误", zap.Error(err))
		return res
	}
	data := string(bytedata)
	data = strings.ReplaceAll(data, " ", "")
	fmt.Println(data)
	var dataJson []map[string]any
	json.Unmarshal([]byte(data), &dataJson)
	fmt.Printf("\n%v", dataJson)
	return bytedata
}

// GetInfo 获取详细信息
func (v *VideoApi) GetInfo(c *gin.Context) any {
	ignores := map[string]any{
		"playarr":    0,
		"playarr_fs": 0,
		"playarr_uk": 0,
		"playarr_bj": 0,
		"playarr_wj": 0,
	}
	defaultRes := map[string]any{}
	keyword := c.Query("url")
	resp, err := http.Get("http://v.58hda.com:8077/ne2/s" + keyword + ".js")
	if err != nil {
		v.Logger.Error("request info err", zap.Error(err))
		return defaultRes
	}
	defer resp.Body.Close()
	allbyte, err := io.ReadAll(resp.Body)
	if err != nil {
		v.Logger.Error("read all err", zap.Error(err))
		return defaultRes
	}
	redisClint := library.NewRedis()
	redisVal, err := redisClint.Get(c, fmt.Sprintf(consts.REDIS_USER_TELEPLAY, v.URK, keyword)).Result()
	if err != nil {
		v.Logger.Error("read all err", zap.Error(err))
	}
	reg := regexp.MustCompile(`(\w+)\[(\d+)\]=\"(.*)\"`)
	res := make(map[string]map[string]any, 0)
	allline := strings.Split(string(allbyte), ";")
	for _, v := range allline {
		if !strings.Contains(v, "http") && !strings.Contains(v, "https") {
			continue
		}
		matches := reg.FindStringSubmatch(v)
		if len(matches) == 4 {
			key := fmt.Sprintf("%s", matches[1]) // 不包含数组索引
			value := matches[3]
			// 忽略一些不需要的线路
			if _, ok := ignores[key]; ok {
				continue
			}
			v, ok := res[key]
			if !ok {
				v = make(map[string]any, 0)
				v["list"] = make([]map[string]string, 0)
				v["info"] = map[string]string{
					"play": "0",
				}
				res[key] = v
			}
			list := v["list"].([]map[string]string)
			item := make(map[string]string)
			content := strings.Split(value, ",")
			item["url"] = content[0]
			item["name"] = content[len(content)-1]
			item["play"] = "0"
			item["startTime"] = "0"
			list = append(list, item)
			v["list"] = list
			if redisVal != "" {
				splits := strings.Split(redisVal, ",")
				if len(splits) > 1 && splits[0] == key && splits[1] == item["name"] {
					item["play"] = "1"
					infoMap := v["info"].(map[string]string)
					infoMap["play"] = "1"
					if len(splits) >= 2 {
						item["startTime"] = splits[2]
					}
				}
			}
		}
	}
	return res
}

// Star 收藏
func (v *VideoApi) Star(c *gin.Context) any {
	var param video.StarParam
	c.ShouldBindJSON(&param)
	if !param.Check() {
		return errcode.PARAM_ERR
	}
	starDB := db.NewStarDB(v.MysqlDB)
	res := &Respone{}
	er, b := starDB.Exists(v.User.Id, param.ID)
	if er != nil {
		return er
	}
	if b {
		er, _ = starDB.DeletedByTvId(v.User.Id, param.ID)
		if er != nil {
			return er
		}
		res.Msg = "取消收藏成功"
		return res.OK2()
	}
	err, id := utils.ID()
	if err != nil {
		return errcode.ApiCustom(err.Error())
	}
	now := time.Now().Unix()
	rerr := starDB.Create(&db.Star{
		Id:         id,
		Name:       param.Name,
		Url:        param.Url,
		TvID:       param.ID,
		UserID:     v.User.Id,
		Createtime: now,
		Updatetime: now,
	})
	if rerr != nil {
		return rerr
	}
	res.Msg = "收藏成功"
	return res.OK2()
}

// PlayRecord 播放记录
func (v *VideoApi) PlayRecord(c *gin.Context) any {
	res := &Respone{}
	var param video.PlayRecordParam
	c.ShouldBindQuery(&param)
	res.OK(map[string]string{
		"msg": "OK",
	})
	redisKey := fmt.Sprintf(consts.REDIS_USER_TELEPLAY, v.URK, param.Teleplay)
	userInfoKey := fmt.Sprintf(consts.REDIS_USER_INFO, v.URK)
	redisClient := library.NewRedis()
	if param.StartTime == "" {
		playRecord, _ := redisClient.Get(c, redisKey).Result()
		if playRecord != "" {
			splits := strings.Split(playRecord, ",")
			// 集数与渠道一致替换否则就不需要替换
			if len(splits) >= 2 && splits[0] == param.Index && splits[1] == param.Name {
				param.StartTime = splits[2]
			}
		}
	}
	statice := redisClient.SetEX(c, redisKey, fmt.Sprintf("%s,%s,%s", param.Index, param.Name, param.StartTime), time.Hour*24*60)
	userInfo := redisClient.Get(c, userInfoKey)
	userInfoStr, err := userInfo.Result()
	if err != nil {
		v.Logger.Sugar().Error(err)
	}
	var resMap map[string]any
	json.Unmarshal([]byte(userInfoStr), &resMap)
	if resMap == nil {
		resMap = make(map[string]any, 0)
		resMap["teleplays"] = []any{}
	}
	teleplays := resMap["teleplays"].([]any)
	teleplays = append(teleplays, param.Teleplay)
	resMap["teleplays"] = teleplays
	resMapByte, err := json.Marshal(resMap)
	if err != nil {
		v.Logger.Sugar().Error(err)
	}
	redisClient.Set(c, userInfoKey, string(resMapByte), time.Hour*24*180)
	if statice.Err() != nil {
		v.Logger.Error("redis play record err", zap.Error(statice.Err()))
	}
	return res
}

// StarList 收藏列表
func (v *VideoApi) StarList(c *gin.Context) any {
	starDB := db.NewStarDB(v.MysqlDB)
	errcode, stars := starDB.ListAll(v.User.Id)
	if errcode != nil {
		return errcode
	}
	r := &Respone{}
	return r.OK(stars)
}

// Download 下载视频
func (v *VideoApi) Download(c *gin.Context) any {
	var param struct {
		URL string `json:"url" form:"url"`
	}
	c.ShouldBindQuery(&param)
	if len(param.URL) <= 0 {
		return errcode.PARAM_ERR
	}
	library.Logger.Sugar().Info(param.URL)
	// m3u8 := ufile.NewM3U8(param.URL, "", func(node, total int) {
	// 	// 进度回传
	// 	progress := strconv.Itoa((node / total) * 100)
	// 	v.Logger.Info("下载进度", zap.String("progress", progress))
	// 	c.Writer.Write([]byte(progress))
	// })

	// mixed := m3u8.CheckMixed()
	// if mixed != "" {
	// 	m3u8.SetSourcePath(mixed)
	// }
	// b, targetPath := m3u8.ToMP4()
	if true {
		file, err := os.Open("/Users/huxuekuo/Downloads/1741748442442668000/1741748679838553000.mp4")
		if err != nil {
			v.Logger.Warn("打开mp4文件失败", zap.Error(err))
		} else {
			return file
		}
	}
	return errcode.VIDEO_DOWNLOAD_ERR
}

// MonitorUpdate 监听更新列表
func (v *VideoApi) MonitorUpdate(c *gin.Context) any {

	return nil
}
