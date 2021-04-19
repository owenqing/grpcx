package cronx

import "github.com/robfig/cron"

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
	AddTask(taskkey, cronexpr string, f func())
	// 启动
	RunAll()
	// 停止
	Stop(taskkey string)
}
