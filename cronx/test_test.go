package cronx

import (
	"fmt"
	"testing"
)

func TestStartCronJob(t *testing.T) {
	ch := make(chan struct{}, 1)
	cronTask := NewCronTasks()
	cronTask.RegisteTask("test1", "*/5 * * * * ?", func() {
		fmt.Println("定时任务1")
	})
	cronTask.RegisteTask("test2", "*/2 * * * * ?", func() {
		fmt.Println("定时任务2")
	})
	cronTask.RunAll()
	<-ch
}
