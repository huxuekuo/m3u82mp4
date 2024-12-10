package systemser

import (
	"encoding/json"
	"m3u82mp4/consts"
	"m3u82mp4/consts/errcode"
	"m3u82mp4/library"
	"m3u82mp4/model/common"
	"m3u82mp4/model/system"
	"time"

	"github.com/gin-gonic/gin"
)

type ProbeState struct {
	State int   `json:"s"`
	Time  int64 `json:"t"`
}

// 查询现有探针
func QueryProbe(c *gin.Context, system *system.VideoURLProbeParam) any {
	redisClint := library.NewRedis()
	urls, _ := redisClint.HGetAll(c, consts.REDIS_VIDEO_QUERY_PROBE).Result()
	array := make([]map[string]any, 0)
	for key, value := range urls {
		// 异常状态
		var s ProbeState
		// 正常状态
		_ = json.Unmarshal([]byte(value), &s)
		item := map[string]any{
			"url":   key,
			"state": s,
		}
		array = append(array, item)
	}
	res := &common.Respone{}
	return res.OK(array)
}

// 添加探针
func AddProbe(c *gin.Context, system *system.VideoURLProbeParam) any {
	redisClint := library.NewRedis()
	s := &ProbeState{
		Time: time.Now().Unix(),
	}
	sStr, _ := json.Marshal(s)
	redisClint.HSet(c, consts.REDIS_VIDEO_QUERY_PROBE, []string{system.URL, string(sStr)})
	res := &common.Respone{}
	return res.OK2()
}

// 获取有效地址
func GetURL(c *gin.Context) string {
	redisClint := library.NewRedis()
	urls, _ := redisClint.HGetAll(c, consts.REDIS_VIDEO_QUERY_PROBE).Result()
	array := make([]map[string]any, 0)
	for key, value := range urls {
		// 异常状态
		var s ProbeState
		// 正常状态
		_ = json.Unmarshal([]byte(value), &s)
		if s.State == 1 {
			item := map[string]any{
				"url":   key,
				"state": s,
			}
			array = append(array, item)
		}
	}
	if len(array) <= 0 {
		library.Logger.Sugar().Error("无可用地址!")
		return ""
	}
	return array[0]["url"].(string)
}

func GetURLApi(c *gin.Context, system *system.VideoURLProbeParam) any {
	res := &common.Respone{}
	return res.OK(GetURL(c))
}

func Deleted(c *gin.Context, system *system.VideoURLProbeParam) any {
	redisClint := library.NewRedis()
	success, err := redisClint.HDel(c, consts.REDIS_VIDEO_QUERY_PROBE, system.URL).Result()
	res := &common.Respone{}
	if err != nil {
		return errcode.ApiCustom("读取信息失败")
	}
	return res.OK(success)
}
