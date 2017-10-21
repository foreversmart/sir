package task

import (
	"fmt"
	"sir/lib/psutil"
	"sir/models"
	"time"
)

type TaskRuntime struct {
	*models.Task

	TaskLogSignal   chan bool
	TaskStateSignal chan bool
}

func NewTaskRuntime(task *models.Task) *TaskRuntime {
	taskRuntime := &TaskRuntime{
		Task:            task,
		TaskLogSignal:   make(chan bool),
		TaskStateSignal: make(chan bool),
	}

	return taskRuntime

}

func (t *TaskRuntime) TaskLog() {
}

func (t *TaskRuntime) TaskStateFunc() {
	for {
		select {
		// kill
		case <-t.TaskStateSignal:
			return

		default:
			state, err := psutil.TaskState(t.Pid)
			if err != nil {
				fmt.Println(err)
			}

			if state == nil {
				state = &models.TaskState{
					Pid: t.Pid,
				}
			}

			// update task state
			t.TaskState = state

			time.Sleep(time.Second)

		}
	}

}

func (t *TaskRuntime) Stop() {
	t.TaskLogSignal <- true
	t.TaskStateSignal <- true
}

func (t *TaskRuntime) Run() {
	go t.TaskLog()

	go t.TaskStateFunc()
}
