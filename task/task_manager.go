package task

import (
	"fmt"
	"os"
	"path"
	"sir/models"
	"sync"
)

type TaskManager struct {
	mutex sync.Mutex
	Tasks map[string]*TaskRuntime

	Workspace string
}

type TaskFlow struct {
	File     *os.File
	FilePath string
}

func NewTaskManager(workspace string) *TaskManager {
	return &TaskManager{
		Tasks:     make(map[string]*TaskRuntime),
		Workspace: workspace,
	}
}

func (t *TaskManager) StopTask(name string) (err error) {
	if task, ok := t.Tasks[name]; ok {
		return task.Stop()

	}

	return fmt.Errorf("%s %s", name, " task not found or already stopped")
}

func (t *TaskManager) StartTask(task *models.Task) (err error) {
	if t.IsTaskExist(task.Name) {
		return
	}

	if task.TaskState == nil {
		task.TaskState = &models.TaskState{}
	}

	// set process flow
	flows, err := t.GenerateTaskFlow(task.Name)
	if err != nil {
		return err
	}

	// task flow
	task.TaskFlows = &models.TaskFlows{
		StdIn:      flows[0].File,
		StdOut:     flows[1].File,
		StdOutPath: flows[1].FilePath,
		StdErr:     flows[2].File,
		StdErrPath: flows[2].FilePath,
	}

	// exec task self func
	taskRuntime := NewTaskRuntime(task)

	// start task
	err = taskRuntime.Start()
	if err != nil {
		return err
	}

	taskRuntime.Run()

	// add task
	t.AddTask(taskRuntime)
	return nil
}

func (t *TaskManager) IsTaskExist(name string) (ok bool) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	_, ok = t.Tasks[name]
	return
}

func (t *TaskManager) AddTask(task *TaskRuntime) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	t.Tasks[task.Name] = task

}

func (t *TaskManager) RemoveTask(task *TaskRuntime) (err error) {
	err = t.StopTask(task.Name)
	if err != nil {
		return
	}

	t.mutex.Lock()
	defer t.mutex.Unlock()

	delete(t.Tasks, task.Name)
	return
}

func (t *TaskManager) GenerateTaskFlow(name string) (flows []*TaskFlow, err error) {
	flows = make([]*TaskFlow, 3)

	// TODO handle error
	os.MkdirAll(t.Workspace, 0700)

	flowPath := path.Join(t.Workspace, name+".temp.in")

	flows[0] = &TaskFlow{}
	flows[0].FilePath = flowPath
	flows[0].File, err = os.OpenFile(flowPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return flows, err
	}

	flowPath = path.Join(t.Workspace, name+".temp.stdout")

	flows[1] = &TaskFlow{}
	flows[1].FilePath = flowPath
	flows[1].File, err = os.OpenFile(flowPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return flows, err
	}

	flowPath = path.Join(t.Workspace, name+".temp.stderr")

	flows[2] = &TaskFlow{}
	flows[2].FilePath = flowPath
	flows[2].File, err = os.OpenFile(flowPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return flows, err
	}

	return
}
