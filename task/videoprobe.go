package task

import (
	"context"
	"encoding/json"
	"m3u82mp4/consts"
	"m3u82mp4/library"
	systemser "m3u82mp4/model/service/system"
	"net/http"
	"time"
)

func VideoProbe() {
	library.Logger.Sugar().Info("videoProbe-start")
	c := context.Background()
	redisClint := library.NewRedis()
	tr := &http.Transport{
		ResponseHeaderTimeout: 5 * time.Second,
	}
	client := &http.Client{
		Transport: tr,
	}
	urls, err := redisClint.HGetAll(c, consts.REDIS_VIDEO_QUERY_PROBE).Result()
	if err != nil {
		library.Logger.Sugar().Info("videoProbe-redis-error-end")
		return
	}
	if len(urls) <= 0 {
		library.Logger.Sugar().Info("videoProbe-url-empty-end")
		return
	}
	for key, value := range urls {
		_, err := client.Get(key)
		state := 0
		if err == nil {
			state = 1
		}
		// 异常状态
		var s systemser.ProbeState
		// 正常状态
		_ = json.Unmarshal([]byte(value), &s)
		s.State = state
		s.Time = time.Now().Unix()
		sStr, _ := json.Marshal(s)
		redisClint.HSet(c, consts.REDIS_VIDEO_QUERY_PROBE, []string{key, string(sStr)})
	}
	library.Logger.Sugar().Info("videoProbe-end")
}
