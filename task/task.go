package task

import "sir/models"

type TaskRuntime struct {
	*models.Task

	TaskLogSignal   chan bool
	TaskStateSignal chan bool
}

func NewTaskRuntime(task *models.Task) *TaskRuntime {
	taskRuntime := TaskRuntime{
		task,
		TaskLogSignal:   make(chan bool),
		TaskStateSignal: make(chan bool),
	}

	return taskRuntime

}

func (t *TaskRuntime) TaskLog() {
}

func (t *TaskRuntime) TaskState() {

}

func (t *TaskRuntime) Run() {
	go t.TaskLog()

	go t.TaskState()
}
