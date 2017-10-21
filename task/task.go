package task

import (
	"fmt"
	"sir/lib/psutil"
	"sir/models"
	"time"

	"log"
	"os"
	"os/signal"
	"sir/lib/config"
	"sir/models"
	"syscall"

	readline "github.com/jprichardson/readline-go"
	"github.com/natefinch/lumberjack"
)

type TaskRuntime struct {
	*models.Task

	TaskStdLogSignal   chan bool
	TaskErrorLogSignal chan bool
	TaskStateSignal    chan bool
}

func NewTaskRuntime(task *models.Task) *TaskRuntime {
	taskRuntime := &TaskRuntime{
		Task:               task,
		TaskStdLogSignal:   make(chan bool),
		TaskErrorLogSignal: make(chan bool),
		TaskStateSignal:    make(chan bool),
	}

	return taskRuntime
}

func (t *TaskRuntime) TaskLog() {

	// deal with std log
	go func() {
		logger := &lumberjack.Logger{
			Filename:   config.DefaultLogPath + "/log.log",
			MaxSize:    10, // megabytes
			MaxBackups: 0,
			MaxAge:     0, //days
		}

		log.SetOutput(logger)

		readline.ReadLine(t.Task.TaskFlows.StdOut, func(line string) {

			grSignal := make(chan bool)

			select {
			case <-t.TaskStdLogSignal:
				grSignal <- true
				return
			}

			logger.Write([]byte(line))

			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGHUP)

			go func() {
				for {

					select {
					case <-grSignal:
						return
					}

					<-c
					logger.Rotate()
				}
			}()
		})
	}()

	// deal with error log
	go func() {
		logger := &lumberjack.Logger{
			Filename:   config.DefaultLogPath + "/error.log",
			MaxSize:    10, // megabytes
			MaxBackups: 0,
			MaxAge:     0, //days
		}

		log.SetOutput(logger)

		readline.ReadLine(t.Task.TaskFlows.StdErr, func(line string) {

			grSignal := make(chan bool)

			select {
			case <-t.TaskErrorLogSignal:
				grSignal <- true
				return
			}

			logger.Write([]byte(line))

			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGHUP)

			go func() {
				for {

					select {
					case <-grSignal:
						return
					}

					<-c
					logger.Rotate()
				}
			}()
		})
	}()
}

func (t *TaskRuntime) TaskStateFunc() {
	for {
		select {
		// killz
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
	t.TaskStdLogSignal <- true
	t.TaskErrorLogSignal <- true
	t.TaskStateSignal <- true
}

func (t *TaskRuntime) Run() {
	go t.TaskLog()

	go t.TaskStateFunc()
}
