package models

import (
	"time"

	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
)

type Task struct {
	*TaskState

	*TaskConfig
}

type TaskState struct {
	Pid        int32                   `json:"pid"`
	CpuPercent float64                 `json:"cpu"`
	Mem        uint64                  `json:"mem"` // Kb
	MemPercent float32                 `json:"mem"`
	Disk       int                     `json:"disk"`
	IoCounter  *process.IOCountersStat `json:"io_counter"`
	Net        *net.IOCountersStat     `json:"net"`
	Load       float64                 `json:"load"`
	Stat       string                  `json:"stat"`
	UpTime     int64                   `json:"up_time"` // timestamp(msec)
}

type TaskConfig struct {
	Name            string            `json:"name"`
	Cmd             string            `json:"cmd"`
	Watch           bool              `json:"watch"`
	WatchDir        string            `json:"watch_dir"`
	Env             map[string]string `json:"env"`
	Workspace       string            `json:"workspace"`
	User            string            `json:"user"`
	Group           string            `json:"group"`
	Priority        int               `json:"priority"` // higher is lower
	AutoRestart     bool              `json:"autorestart"`
	AutoStart       bool              `json:"auto_start"`
	RestartInterval int               `json:"restart_interval"` // seconds
	Rules           []*RuleConfig     `json:"rules"`
	LogConfigs      []*LogConfig      `json:"log_configs"`

	//
	CTime        time.Time `json:"c_time"`
	RestartCount int       `json:"restart_count"`
}

type RuleConfig struct {
	Type      string  `json:"type"`
	Threshold float64 `json:"threshold"`
}

type LogConfig struct {
	ErrLogPath string `json:"err_log_path"`
	StdLogPath string `json:"std_log_path"`
	RotateType string `json:"rotate_type"`
	Limit      int    `json:"limit"`
}
