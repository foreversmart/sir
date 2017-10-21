package models

import (
	"time"

	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
	"os"
	"strings"
)

type Task struct {
	*TaskState

	*TaskFlows

	*TaskConfig
}

func (t *Task) ParseCmd() (cmd string, args []string) {
	items := strings.Split(t.Cmd, " ")
	cmd = items[0]
	if len(items) > 1 {
		args = items[1:]
	}

	return
}

type TaskFlows struct {
	StdIn  *os.File
	StdOut *os.File
	StdErr *os.File
}

type TaskState struct {
	Pid        int                     `json:"pid"`
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
	Name            string        `json:"name"`
	Cmd             string        `json:"cmd"`
	Watch           bool          `json:"watch"`
	WatchDir        string        `json:"watch_dir"`
	Env             []string      `json:"env"` // key=value
	Workspace       string        `json:"workspace"`
	User            string        `json:"user"`
	Group           string        `json:"group"`
	Priority        int           `json:"priority"` // higher is lower
	AutoRestart     bool          `json:"autorestart"`
	AutoStart       bool          `json:"auto_start"`
	RestartInterval int           `json:"restart_interval"` // seconds
	Rules           []*RuleConfig `json:"rules"`
	LogConfigs      *LogConfig    `json:"log_configs"`

	//
	CTime        time.Time `json:"c_time"`
	RestartCount int       `json:"restart_count"`
}

func (conf *TaskConfig) IsValid() bool {
	return conf.Name != ""
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
