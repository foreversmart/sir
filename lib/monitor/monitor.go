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

func PushMonitorData(state *models.TaskState) {
	reg = prometheus.NewRegistry()

	// init the cpu data
	CpuData := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "process",
			Subsystem: "CPU",
			Name:      "cpu_percent",
			Help:      "docstring",
		},
		[]string{"pid", "unit"},
	)

	CpuData.WithLabelValues(strconv.Itoa(state.Pid), "%").Set(state.CpuPercent)

	// init the mem usage data
	memUsageData := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "process",
			Subsystem: "mem",
			Name:      "mem_usage",
			Help:      "docstring",
		},
		[]string{"pid", "unit"},
	)
	memUsageData.WithLabelValues(strconv.Itoa(state.Pid), "Kb").Set(float64(state.Mem))

	// init the mem percent data
	memPercentData := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "process",
			Subsystem: "mem",
			Name:      "mem_percent",
			Help:      "docstring",
		},
		[]string{"pid", "unit"},
	)
	memPercentData.WithLabelValues(strconv.Itoa(state.Pid), "%").Set(float64(state.MemPercent))

	// register the data
	reg.MustRegister(CpuData)
	reg.MustRegister(memUsageData)
	reg.MustRegister(memPercentData)

	http.Handle("/", handler())
	http.ListenAndServe(":9091", nil)
}
