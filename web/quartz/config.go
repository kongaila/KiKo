package quartz

import (
	"github.com/robfig/cron"
)

const (
	spec30min = "0 0/30 * * * ?"
	FM12      = "0/5 0 12 * * ?"
)

// 定时任务配置
func Config() {
	c := cron.New()
	_ = c.AddFunc(spec30min, top)
	_ = c.AddFunc(FM12, age)
	go c.Start()
}
