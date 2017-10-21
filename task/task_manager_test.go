package task

import (
	"github.com/stretchr/testify/assert"
	"sir/models"
	"testing"
)

func TestTaskManager_Start(t *testing.T) {
	assertion := assert.New(t)

	manager := NewTaskManager("/Users/hong/projects/src/sir/")

	task := &models.Task{
		&models.TaskState{},
		&models.TaskConfig{
			Name: "test",
			Cmd:  "/sbin/ping 114.114.114.114",
		},
	}
	err := manager.Start(task)
	assertion.Nil(err)
	assertion.Equal(1, len(manager.Tasks))
}
