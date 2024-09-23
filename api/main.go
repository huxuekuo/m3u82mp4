package main

import (
	"encoding/json"
	"fmt"
	"io"
	"m3u82mp4/consts"
	"m3u82mp4/library"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"go.uber.org/zap"
)

var logger *zap.Logger

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. [必须]接受指定域的请求，可以使用*不加以限制，但不安全
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		// 2. [必须]设置服务器支持的所有跨域请求的方法
		w.Header().Set("Access-Control-Allow-Methods", "POST,GET,PUT,DELETE,OPTIONS")
		// 3. [可选]服务器支持的所有头信息字段，不限于浏览器在"预检"中请求的字段
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Content-Length,Token")
		// 4. [可选]设置XMLHttpRequest的响应对象能拿到的额外字段
		w.Header().Set("Access-Control-Expose-Headers", "Access-Control-Allow-Headers,Token")
		// 5. [可选]是否允许后续请求携带认证信息Cookir，该值只能是true，不需要则不设置
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		next.ServeHTTP(w, r)
	})
}

// 119.29.226.140:13457 172.247.47.125
func query(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("key")
	other_kkk217 := "http://xdm530.com"
	w.Header().Set("Content-Type", "application/json")
	// 第一次 URL 编码
	encodedStr := url.QueryEscape(other_kkk217)

	// 将编码后的字符串中的 '%' 替换为 '%25'
	encodedStr = strings.Replace(encodedStr, "%", "%25", -1)

	// 第二次 URL 编码
	doubleEncodedStr := url.QueryEscape(encodedStr)
	resp, err := http.Get(fmt.Sprintf("http://119.29.226.140:13457/ssszz.php?top=10&q=%s&other_kkk217=%s&dect=0", url.QueryEscape(key), doubleEncodedStr))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	bytedata, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	w.Write(bytedata)
}

func getInfoV2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	keyword := r.FormValue("url")
	resp, err := http.Get("http://v.58hda.com:8077/ne2/s" + keyword + ".js")
	if err != nil {
		logger.Error("request info err", zap.Error(err))
		return
	}
	defer resp.Body.Close()
	allbyte, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("read all err", zap.Error(err))
		return
	}
	redisClint := library.NewRedis()
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
			list = append(list, item)
			v["list"] = list
			redisKey := fmt.Sprintf(consts.REDIS_USER_TELEPLAY, keyword)
			RedisRes := redisClint.Get(r.Context(), redisKey)
			if RedisRes.Err() != nil {
				logger.Sugar().Error(RedisRes.Err())
				continue
			}
			val, _ := RedisRes.Result()
			splits := strings.Split(val, ",")
			if len(splits) > 1 && splits[0] == key && splits[1] == item["name"] {
				item["play"] = "1"
				infoMap := v["info"].(map[string]string)
				infoMap["play"] = "1"
			}
		}
	}

	delete(res, "playarr")
	json.NewEncoder(w).Encode(res)
}

// 播放记录
func PlayRecord(w http.ResponseWriter, r *http.Request) {
	index := r.FormValue("index")
	teleplay := r.FormValue("teleplay")
	name := r.FormValue("name")
	startTime := r.FormValue("startTime")
	uid := 1
	redisKey := fmt.Sprintf(consts.REDIS_USER_TELEPLAY, teleplay)
	userInfoKey := fmt.Sprintf(consts.REDIS_USER_INFO, uid)
	redisClient := library.NewRedis()
	logger.Sugar().Info(redisKey)
	statice := redisClient.SetEX(r.Context(), redisKey, fmt.Sprintf("%s,%s,%s", index, name, startTime), time.Hour*24*60)
	userInfo := redisClient.Get(r.Context(), userInfoKey)
	userInfoStr, err := userInfo.Result()
	if err != nil {
		logger.Sugar().Error(err)
	}
	var resMap map[string]any
	json.Unmarshal([]byte(userInfoStr), &resMap)
	if resMap == nil {
		resMap = make(map[string]any, 0)
		resMap["teleplays"] = make([]string, 0)
	}
	teleplays := resMap["teleplays"].([]string)
	teleplays = append(teleplays, teleplay)
	resMap["teleplays"] = teleplays
	resMapByte, err := json.Marshal(resMap)
	if err != nil {
		logger.Sugar().Error(err)
	}
	redisClient.Set(r.Context(), userInfoKey, string(resMapByte), time.Hour*24*180)
	if statice.Err() != nil {
		logger.Error("redis play record err", zap.Error(statice.Err()))
	}
	res := map[string]string{
		"msg": "OK",
	}
	json.NewEncoder(w).Encode(res)
}

// 历史观看
func playStarTime(w http.ResponseWriter, r *http.Request) {}

func main() {
	logger, _ = zap.NewProduction()
	http.Handle("/query", corsMiddleware(http.HandlerFunc(query)))
	http.Handle("/getInfoV2", corsMiddleware(http.HandlerFunc(getInfoV2)))
	http.Handle("/playRecord", corsMiddleware(http.HandlerFunc(PlayRecord)))
	// http.Handle("/playStarTime",corsMiddleware(http.HandlerFunc()))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		logger.Error("run server error", zap.Error(err))
	}
}
