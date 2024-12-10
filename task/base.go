package task

import "github.com/robfig/cron/v3"

func InitTask() {
	c := cron.New()
	c.AddFunc("@every 5m", VideoProbe)
	c.Start()
}
