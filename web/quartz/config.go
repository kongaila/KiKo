package quartz

import (
	"github.com/robfig/cron"
)

const (
	spec5min = "0 0/5 * * * ?"
	spec10s  = "* * */12 * * ?"
)

// 定时任务配置
func Config() {
	c := cron.New()
	_ = c.AddFunc(spec5min, top)
	_ = c.AddFunc(spec10s, age)
	go c.Start()
}
