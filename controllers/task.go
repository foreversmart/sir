package controllers

import (
	"encoding/json"
	"sir/lib/config"
	"sir/models"

	"sir/task"

	"github.com/astaxie/beego"
)

var TaskManager *task.TaskManager

func init() {
	TaskManager = task.NewTaskManager(config.AppHome + beego.AppConfig.String("workspace"))
}

type TaskController struct {
	beego.Controller
}

// @router /task/add [post]
func (task *TaskController) Add() {
	var ob models.TaskConfig
	json.Unmarshal(task.Ctx.Input.RequestBody, &ob)

	config.CreateTaskConfig(&ob)
	// TODO handle error

	task.Data["json"] = map[string]interface{}{"data": ob}
	task.ServeJSON()
}

// @router /task/:name [delete]
func (task *TaskController) Remove() {
	taskname := task.Ctx.Input.Param(":name")
	err := config.DeleteTaskConfig(taskname)
	if err != nil {
		beego.Error(err)
	}

	// update task runtime
	err = TaskManager.StopTask(taskname)
	if err != nil {
		beego.Error(err)
	}

	task.ServeJSON()
}

// @router /task/:name [put]
func (task *TaskController) Update() {
	taskname := task.Ctx.Input.Param(":name")
	// TODO just open editor

	task.Data["json"] = map[string]interface{}{"data": config.GetTaskConfigFilePath(taskname)}

	task.ServeJSON()
}

// @router /task/:name/rename [put]
func (task *TaskController) Rename() {
}

// @router /task/:name [get]
func (task *TaskController) Show() {
	taskname := task.Ctx.Input.Param(":name")

	tasks := TaskManager.Tasks
	if t, ok := tasks[taskname]; ok {
		task.Data["json"] = map[string]interface{}{"data": *t.Task}
	} else {
		taskConfig, err := config.GetTaskConfig(taskname)
		if err != nil {
			beego.Error("config.GetTaskConfig", err)
			return
		}
		task.Data["json"] = map[string]interface{}{"data": *taskConfig}
	}

	task.ServeJSON()
}

// @router /task [get]
func (task *TaskController) List() {
	// beego.Error("list")

	runningTasks := TaskManager.Tasks

	taskConfigs := config.ListAllTaskConfigs()

	tasks := make([]models.Task, 0, 10)
	for _, c := range taskConfigs {

		if t, ok := runningTasks[c.Name]; ok {
			tasks = append(tasks, *(t.Task))
		} else {
			temp := c
			tasks = append(tasks, models.Task{
				TaskConfig: &temp,
			})
		}
	}

	task.Data["json"] = map[string][]models.Task{
		"data": tasks,
	}
	task.ServeJSON()
}
