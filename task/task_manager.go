package task

import (
	"fmt"
	"os"
	"os/user"
	"path"
	"sir/models"
	"sync"
	"syscall"
)

type TaskManager struct {
	mutex sync.Mutex
	Tasks []*models.Task

	Workspace string
}

func NewTaskManager(workspace string) *TaskManager {
	return &TaskManager{
		Tasks:     make([]*models.Task, 0, 5),
		Workspace: workspace,
	}
}

func (t *TaskManager) Start(task *models.Task) {
	// set work space
	var (
		workspace string
	)

	dir, _ := os.Getwd()
	workspace = dir
	if task.Workspace != "" {
		workspace = task
	}

	// set env
	env := os.Environ()
	env = append(env, task.Env...)

	// set process flow
	flows, err := t.GenerateTaskFlow(task.Name)
	if err != nil {
		return err
	}

	// set uid
	attr := syscall.SysProcAttr{}
	attr.Credential = &syscall.Credential{}
	if task.User != "" {
		taskUser, err := user.Lookup(task.User)
		if err != nil {
			return err
		}

		attr.Credential.Uid = taskUser.Uid
	}

	// set group
	if task.Group != "" {
		taskGroup, err := user.LookupGroup(task.Group)
		if err != nil {
			return err
		}

		attr.Credential.Gid = taskGroup.Gid
	}

	// start task
	procAttrs := os.ProcAttr{Dir: workspace, Env: env, Files: flows, Sys: &attr}
	process, err := os.StartProcess(task.Cmd, []string{}, &procAttrs)
	if err != nil {
		return fmt.Errorf("can't create process %s: %s", os.Args[0], err)
	}

	task.Pid = process.Pid

	// add task
	t.AddTask(task)
	return nil
}

func (t *TaskManager) AddTask(task models.Task) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	t.Tasks = append(t.Tasks, task)
}

func (t *TaskManager) GenerateTaskFlow(name string) (flows []*os.File, err error) {
	flows = make([]*os.File, 3)
	flows[0], err = os.OpenFile(path.Join(t.Workspace, name+".temp.in"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return flows, err
	}

	flows[1], err = os.OpenFile(path.Join(t.Workspace, name+".temp.stdout"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return flows, err
	}

	flows[2], err = os.OpenFile(path.Join(t.Workspace, name+".temp.stderr"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return flows, err
	}

	return
}
