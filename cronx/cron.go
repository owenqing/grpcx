package cronx

import (
	"sync"

	"github.com/robfig/cron"
)

var (
	cronSync         sync.Once
	cronTaskInstance *CronTasks
)

// 定时任务单元
type cronitem struct {
	expr    string
	fn      func()
	cronjob *cron.Cron // 定时任务
}

type CronTasks struct {
	tasks map[string]*cronitem
}

type CronTasksInterface interface {
	// 新增定时任务
	RegisteTask(taskkey, cronexpr string, f func())
	// 启动
	RunAll()
	// 停止
	Stop(taskkey string)
}
