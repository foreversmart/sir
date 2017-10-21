package mock

import "sir/models"
import "time"

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
