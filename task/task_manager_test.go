package task

import (
	"testing"
	"time"

	"github.com/foreversmart/sir/models"
	"github.com/stretchr/testify/assert"
)

func TestTaskManager_StartTask(t *testing.T) {
	assertion := assert.New(t)

	manager := NewTaskManager("/Users/hong/projects/src/sir/")

	task := &models.Task{
		TaskState: &models.TaskState{},
		TaskConfig: &models.TaskConfig{
			Name: "test",
			Cmd:  "/sbin/ping 114.114.114.114",
		},
	}
	err := manager.StartTask(task)
	assertion.Nil(err)
	assertion.Equal(1, len(manager.Tasks))
}

func TestTaskManager_StopTask(t *testing.T) {
	assertion := assert.New(t)

	manager := NewTaskManager("/Users/hong/projects/src/sir/")

	task := &models.Task{
		TaskState: &models.TaskState{},
		TaskConfig: &models.TaskConfig{
			Name: "test",
			Cmd:  "/sbin/ping 114.114.114.114",
		},
	}
	err := manager.StartTask(task)
	assertion.Nil(err)

	err = manager.StopTask(task.Name)
	assertion.Nil(err)

}

func TestTaskManager_TaskState(t *testing.T) {
	assertion := assert.New(t)

	manager := NewTaskManager("/Users/hong/projects/src/sir/")

	task := &models.Task{
		TaskState: &models.TaskState{},
		TaskConfig: &models.TaskConfig{
			Name: "test",
			Cmd:  "/sbin/ping 114.114.114.114",
		},
	}
	err := manager.StartTask(task)
	assertion.Nil(err)

	time.Sleep(time.Second)

	assertion.Empty(manager.Tasks["test"].TaskState)
}
