package systemser

import (
	"m3u82mp4/consts"
	"m3u82mp4/library"
	"m3u82mp4/model/db"

	"github.com/gin-gonic/gin"
)

type ProbeState struct {
	State int   `json:"s"`
	Time  int64 `json:"t"`
}

var index = 0

// 获取有效地址
func GetURL(c *gin.Context) string {
	redisClint := library.NewRedis()
	invalidUrls, _ := redisClint.LRange(c, consts.REDIS_VIDEO_QUERY_PROBE, 0, -1).Result()
	sysDb := db.NewSystemConfigDB(library.MysqlDB)
	errcode, allUrl := sysDb.QueryList("video_config", "")
	if errcode != nil {
		library.Logger.Sugar().Error("Mysql中无可用地址!")
		return ""
	}
	array := make([]string, 0)
	for _, value := range allUrl {
		ex := false
		// 黑名单url
		for _, exValue := range invalidUrls {
			if exValue == value.Key {
				ex = true
			}
		}
		// 有效地址
		if !ex {
			array = append(array, value.Value)
		}

	}
	if len(array) <= 0 {
		library.Logger.Sugar().Error("过滤后无可用地址!")
		return ""
	}
	// 顺序取
	len := len(array)
	if len <= index {
		index = 0
	}
	resurl := array[index]
	index++
	return resurl
}
