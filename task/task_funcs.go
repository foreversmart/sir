package task

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sir/lib/monitor"
	"sir/lib/psutil"
	"sir/models"
	"syscall"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/jprichardson/readline-go"
	"github.com/natefinch/lumberjack"
)

func (t *TaskRuntime) TaskLog() {

	// deal with std log
	go func() {
		logger := &lumberjack.Logger{
			Filename:   t.Task.TaskConfig.LogConfigs.StdLogPath + "/log.log",
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
			Filename:   t.Task.TaskConfig.LogConfigs.ErrLogPath + "/error.log",
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

			if t.Task.TaskConfig.Monitor {
				monitor.PushMonitorData(state)
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

func (t *TaskRuntime) TaskWatchFunc() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	defer watcher.Close()

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					err = t.Stop()
					if err != nil {
						log.Println("watch stop", err)
					}

					err = t.Start()
					if err != nil {
						log.Println("watch start", err)
					}
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
			case <-t.TaskStateSignal:
				return

			}
		}
	}()

	err = watcher.Add(t.WatchDir)
	if err != nil {
		log.Fatal(err)
	}

}

func (t *TaskRuntime) StopFunc() {
	t.TaskStdLogSignal <- true
	t.TaskErrorLogSignal <- true
	t.TaskStateSignal <- true

	if t.Watch && t.WatchDir != "" {
		t.TaskWatchSignal <- true
	}
}
