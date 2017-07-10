package main

import (
	"os/exec"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	serviceActive = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "service",
		Name:      "up",
		Help:      "Check if the Service is up.",
	})
)

// Exporter implements the prometheus.Collector interface.
type Exporter struct {
	service string
}

func NewExporter(s string) *Exporter {
	return &Exporter{
		service: s,
	}
}

// Describe describes all the registered stats metrics.
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	serviceActive.Describe(ch)
}

// Collect collects all the registered stats metrics.
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	if err := e.collect(); err != nil {
		return
	}
	serviceActive.Collect(ch)
}

func (e *Exporter) collect() error {
	serviceActive.Set(serviceCheck(e.service))

	return nil
}

func serviceCheck(s string) float64 {
	//Command to check if systemd service is active
	var up float64

	cmdName := "/bin/systemctl"
	cmdArgs := []string{"is-active", s}

	cmdOut, _ := exec.Command(cmdName, cmdArgs...).Output()

	isActive := strings.TrimSpace(string(cmdOut))

	if isActive == "active" {
		up = 1
	} else {
		up = 0
	}

	return up
}
