package psutil

import (
	"sir/models"

	"log"

	"github.com/shirou/gopsutil/process"
)

func TaskState(pid int) (taskState *models.TaskState, err error) {
	taskState = &models.TaskState{}

	// inject the pid
	taskState.Pid = pid

	pro, err := process.NewProcess(int32(pid))
	if err != nil {
		log.Printf("process.NewProcess(%d): %v", pid, err)
		return
	}

	//----------------------------------------------------------

	cpuPercent, err := pro.CPUPercent()
	if err != nil {
		log.Printf("pro.CPUPercent(): %v", err)
		return
	}
	// inject the cpu
	taskState.CpuPercent = cpuPercent

	mem, err := pro.MemoryInfo()
	if err != nil {
		log.Printf("pro.MemoryInfo(): %v", err)
		return
	}
	// inject the mem
	taskState.Mem = mem.RSS / 1024

	memPercent, err := pro.MemoryPercent()
	if err != nil {
		log.Printf("pro.MemoryPercent(): %v", err)
		return
	}
	// inject the mem percent
	taskState.MemPercent = memPercent

	net, errS := pro.NetIOCounters(false)
	if errS != nil {
		log.Printf("pro.NetIOCounters(false): %v", errS)
		// ignore it on mac
	}
	// inject the net BytesSent and BytesRecv
	if len(net) > 0 {
		taskState.Net = &net[0]
	}

	io, err := pro.IOCounters()
	if err != nil {
		log.Printf("pro.IOCounters(): %v", err)
	}

	// inject the io counter
	taskState.IoCounter = io

	stat, err := pro.Status()
	if err != nil {
		log.Printf("pro.Status(): %v", err)
	}
	// inject the stat
	taskState.Stat = stat

	uptime, err := pro.CreateTime()
	if err != nil {
		log.Printf("pro.CreateTime(): %v", err)
		return
	}
	// inject the uptime
	taskState.UpTime = uptime

	return
}
