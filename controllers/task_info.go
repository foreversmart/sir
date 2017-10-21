package controllers

// @router /task/:name/statistics [get]
func (task *TaskController) Statistics() {

}

// @router /task/:name/log [get]
func (task *TaskController) Log() {
	taskname := task.Ctx.Input.Param(":name")

	t := TaskManager.Tasks[taskname]
	task.Data["json"] = map[string]interface{}{
		"err": t.StdErrPath,
		"std": t.StdOutPath,
	}

	task.ServeJSON()
}

// @router /task/log [get]
func (task *TaskController) AllLog() {
	tasks := TaskManager.Tasks

	errLogs := []string{}
	stdLogs := []string{}

	for _, t := range tasks {
		errLogs = append(errLogs, t.StdErrPath)
		stdLogs = append(stdLogs, t.StdOutPath)
	}

	task.Data["json"] = map[string]interface{}{
		"errs": errLogs,
		"stds": stdLogs,
	}

	task.ServeJSON()
}
