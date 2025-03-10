package task

import (
	"context"
	"m3u82mp4/consts"
	"m3u82mp4/library"
	"m3u82mp4/model/db"
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

	sysConfigDb := db.NewSystemConfigDB(library.MysqlDB)
	errcode, urls := sysConfigDb.QueryList("video_config", "")
	if errcode != nil {
		return
	}
	if len(urls) <= 0 {
		library.Logger.Sugar().Info("videoProbe-url-empty-end")
		return
	}
	invalidUrl := []string{}
	for _, value := range urls {
		_, err := client.Get(value.Value)
		if err == nil {
			continue
		}
		invalidUrl = append(invalidUrl, value.Key)
	}
	redisClint.LPush(c, consts.REDIS_VIDEO_QUERY_PROBE, invalidUrl)
	library.Logger.Sugar().Info("videoProbe-end", invalidUrl)
}
