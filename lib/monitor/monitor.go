package monitor

import (
	"bytes"
	"log"
	"net/http"
	"strconv"

	"github.com/foreversmart/sir/models"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	reg *prometheus.Registry

	cpuData        *prometheus.GaugeVec
	memUsageData   *prometheus.GaugeVec
	memPercentData *prometheus.GaugeVec
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

func StartMonitor() {

	reg = prometheus.NewRegistry()

	// init the cpu data
	cpuData = initGuage("process", "CPU", "cpu_percent", []string{"pid", "unit"})

	// init the mem usage data
	memUsageData = initGuage("process", "mem", "mem_usage", []string{"pid", "unit"})

	// init the mem percent data
	memPercentData = initGuage("process", "mem", "mem_percent", []string{"pid", "unit"})

	// register the data
	reg.MustRegister(cpuData)
	reg.MustRegister(memUsageData)
	reg.MustRegister(memPercentData)

	http.Handle("/", handler())
	http.ListenAndServe(":9091", nil)
}

func PushMonitorData(state *models.TaskState) {

	pid := strconv.Itoa(state.Pid)

	cpuData.WithLabelValues(pid, "%").Set(state.CpuPercent)

	memUsageData.WithLabelValues(pid, "Kb").Set(float64(state.Mem))

	memPercentData.WithLabelValues(pid, "%").Set(float64(state.MemPercent))
}
