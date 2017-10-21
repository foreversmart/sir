package mock

import (
	"sir/models"
	"time"

	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
)

func GetTaskConfig() models.TaskConfig {
	return models.TaskConfig{
		Name:            "hackathon",
		Cmd:             "/home/alex/hackathon/server",
		Watch:           true,
		WatchDir:        "/home/alex/hackathon/",
		Env:             map[string]string{"MODE": "dev", "path": "/xxx"},
		Workspace:       "",
		User:            "alex",
		Group:           "alex",
		Priority:        0,
		AutoRestart:     true,
		AutoStart:       true,
		RestartInterval: 60,
		Rules: []*models.RuleConfig{
			&models.RuleConfig{
				Type:      "cpu",
				Threshold: 10,
			},
			&models.RuleConfig{
				Type:      "mem",
				Threshold: 1024,
			},
		},
		LogConfigs: &models.LogConfig{
			ErrLogPath: "/logs/",
			StdLogPath: "/logs/",
			RotateType: "day",
			Limit:      1,
		},
		CTime:        time.Now(),
		RestartCount: 7,
	}
}

func GetTaskState() models.TaskState {
	return models.TaskState{
		Pid:        12342,
		CpuPercent: 15.3,
		Mem:        43,
		MemPercent: 1.5,
		Disk:       540,
		IoCounter: &process.IOCountersStat{
			ReadCount:  12,
			WriteCount: 32,
			ReadBytes:  44,
			WriteBytes: 55,
		},
		Net: &net.IOCountersStat{
			BytesSent:   11,
			BytesRecv:   11,
			PacketsSent: 11,
			PacketsRecv: 11,
		},
		Load:   1.2,
		Stat:   "S+",
		UpTime: 1508577481,
	}
}

func GetTask() models.Task {
	config := GetTaskConfig()
	state := GetTaskState()
	return models.Task{
		TaskConfig: &config,
		TaskState:  &state,
	}
}

func GetTaskList() []models.Task {
	task1 := GetTask()
	task2 := GetTask()
	task2.TaskState = nil
	task2.TaskConfig.Name = "notrunning"
	return []models.Task{
		task2,
		task1,
	}
}
