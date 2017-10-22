package monitor

import (
	"bytes"
	"log"
	"net/http"
	"sir/models"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	reg *prometheus.Registry
)

func handler() http.Handler {

	logBuf := &bytes.Buffer{}

	logger := log.New(logBuf, "", 0)

	return promhttp.HandlerFor(reg, promhttp.HandlerOpts{
		ErrorLog:      logger,
		ErrorHandling: promhttp.ContinueOnError,
	})
}

func initGuage(nameSpace, subSystem, name string, labels []string) *prometheus.GaugeVec {
	return prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: nameSpace,
			Subsystem: subSystem,
			Name:      name,
			Help:      "docstring",
		},
		labels,
	)
}

func PushMonitorData(state *models.TaskState) {

	reg = prometheus.NewRegistry()

	pid := strconv.Itoa(state.Pid)

	// init the cpu data
	cpuData := initGuage("process", "CPU", "cpu_percent", []string{"pid", "unit"})

	cpuData.WithLabelValues(pid, "%").Set(state.CpuPercent)

	// init the mem usage data
	memUsageData := initGuage("process", "mem", "mem_usage", []string{"pid", "unit"})

	memUsageData.WithLabelValues(pid, "Kb").Set(float64(state.Mem))

	// init the mem percent data
	memPercentData := initGuage("process", "mem", "mem_percent", []string{"pid", "unit"})

	memPercentData.WithLabelValues(pid, "%").Set(float64(state.MemPercent))

	// register the data
	reg.MustRegister(cpuData)
	reg.MustRegister(memUsageData)
	reg.MustRegister(memPercentData)

	http.Handle("/", handler())
	http.ListenAndServe(":9091", nil)
}
