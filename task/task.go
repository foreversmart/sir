package task

import (
	"fmt"

	"github.com/foreversmart/sir/models"

	"os"
	"os/user"
	"strconv"
	"syscall"
)

type TaskRuntime struct {
	*models.Task

	TaskStdLogSignal   chan bool
	TaskErrorLogSignal chan bool
	TaskStateSignal    chan bool
	TaskWatchSignal    chan bool
}

func NewTaskRuntime(task *models.Task) *TaskRuntime {
	taskRuntime := &TaskRuntime{
		Task:               task,
		TaskStdLogSignal:   make(chan bool),
		TaskErrorLogSignal: make(chan bool),
		TaskStateSignal:    make(chan bool),
		TaskWatchSignal:    make(chan bool),
	}

	return taskRuntime
}

func (task *TaskRuntime) Start() (err error) {
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
		attr.Credential.NoSetGroups = true
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
		attr.Credential.NoSetGroups = true
	}

	// set files flow
	files := make([]*os.File, 3)
	files[0] = task.TaskFlows.StdIn
	files[1] = task.TaskFlows.StdOut
	files[2] = task.TaskFlows.StdErr

	// start task
	procAttrs := os.ProcAttr{Dir: workspace, Env: env, Files: files, Sys: &attr}

	cmd, args := task.ParseCmd()
	cmdArgs := append([]string{cmd}, args...)
	process, err := os.StartProcess(cmd, cmdArgs, &procAttrs)
	if err != nil {
		return fmt.Errorf("can't create process %s: %v ||||%s", os.Args[0], os.Args, err)
	}

	task.Pid = process.Pid
	return
}

func (t *TaskRuntime) Stop() (err error) {
	process, err := os.FindProcess(t.Pid)
	if err != nil {
		return err
	}

	err = process.Kill()
	if err == nil {
		t.Pid = -1

	}

	return err
}

func (t *TaskRuntime) Run() {
	go t.TaskLog()

	go t.TaskStateFunc()

	if t.Watch && t.WatchDir != "" {
		go t.TaskWatchFunc()
	}
}
