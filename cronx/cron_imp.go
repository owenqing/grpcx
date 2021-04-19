package cronx

import (
	"github.com/robfig/cron"
)

// 创建定时任务
func NewCronTasks() *CronTasks {
	cronSync.Do(func() {
		cronTaskInstance = &CronTasks{tasks: make(map[string]*cronitem)}
	})
	return cronTaskInstance
}

func (c *CronTasks) RegisteTask(taskkey, cronexpr string, f func()) {
	cronitem := &cronitem{expr: cronexpr, fn: f, cronjob: cron.New()}
	cronitem.cronjob.AddFunc(cronitem.expr, cronitem.fn)
	c.tasks[taskkey] = cronitem
}

// 启动所有定时任务
func (c *CronTasks) RunAll() {
	for _, cronx := range c.tasks {
		cronx.cronjob.Start()
	}
}

// 定制定时任务 by Key
func (c *CronTasks) Stop(taskkey string) {
	c.tasks[taskkey].cronjob.Stop()
}
