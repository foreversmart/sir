package task

import (
	"fmt"
	"os"
	"os/user"
	"path"
	"sir/models"
	"strconv"
	"sync"
	"syscall"
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
		process, err := os.FindProcess(task.Pid)
		if err != nil {
			return err
		}

		err = process.Kill()
		if err == nil {
			task.Pid = -1
		}

		return err
	}

	return fmt.Errorf("%s %s", name, " task not found or already stopped")
}

func (t *TaskManager) StartTask(task *models.Task) (err error) {
	if t.IsTaskExist(task.Name) {
		return
	}

	// set work space
	var (
		workspace string
	)

	dir, _ := os.Getwd()
	workspace = dir
	if task.Workspace != "" {
		workspace = task.Workspace
	}

	// set env
	env := os.Environ()
	env = append(env, task.Env...)

	// set process flow
	flows, err := t.GenerateTaskFlow(task.Name)
	if err != nil {
		return err
	}

	// task flow
	task.TaskFlows = &models.TaskFlows{
		StdIn:      flows[0],
		StdOut:     flows[1].File,
		StdOutPath: flows[1].FilePath,
		StdErr:     flows[2].File,
		StdErrPath: flows[2].FilePath,
	}

	// set uid
	attr := syscall.SysProcAttr{}
	if task.User != "" {
		taskUser, err := user.Lookup(task.User)
		if err != nil {
			return err
		}

		if attr.Credential == nil {
			attr.Credential = &syscall.Credential{}
		}

		uitInt, _ := strconv.ParseUint(taskUser.Uid, 32, 10)
		attr.Credential.Uid = uint32(uitInt)
	}

	// set group
	if task.Group != "" {
		taskGroup, err := user.LookupGroup(task.Group)
		if err != nil {
			return err
		}

		if attr.Credential == nil {
			attr.Credential = &syscall.Credential{}
		}

		groupInt, _ := strconv.ParseUint(taskGroup.Gid, 32, 10)
		attr.Credential.Gid = uint32(groupInt)
	}

	// start task
	procAttrs := os.ProcAttr{Dir: workspace, Env: env, Files: flows, Sys: &attr}

	cmd, args := task.ParseCmd()
	cmdArgs := append([]string{cmd}, args...)
	process, err := os.StartProcess(cmd, cmdArgs, &procAttrs)
	if err != nil {
		return fmt.Errorf("can't create process %s: %v ||||%s", os.Args[0], os.Args, err)
	}

	task.Pid = process.Pid

	// exec task self func
	taskRuntime := NewTaskRuntime(task)
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

	flowPath := path.Join(t.Workspace, name+".temp.in")

	flows[0].FilePath = flowPath
	flows[0].File, err = os.OpenFile(flowPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return flows, err
	}

	flowPath = path.Join(t.Workspace, name+".temp.stdout")

	flows[1].FilePath = flowPath
	flows[1].File, err = os.OpenFile(flowPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return flows, err
	}

	flowPath = path.Join(t.Workspace, name+".temp.stderr")

	flows[2].FilePath = flowPath
	flows[2].File, err = os.OpenFile(flowPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return flows, err
	}

	return
}
