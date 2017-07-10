package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/kelseyhightower/envconfig"
	"github.com/prometheus/client_golang/prometheus"
)

// Specification CLI Arguments for Changing exporter behavior
type Specification struct {
	Debug         bool   `default:"false"`
	ListenAddress string `default:":8080"`
	MetricsPath   string `default:"/metrics"`
	//	Service       string `default:"sshd"`
}

var (
	service = flag.String("s", "", "Name of the service you wish to monitor")
)

func main() {
	flag.Parse()
	var s Specification
	err := envconfig.Process("", &s)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Starting Exporter...")

	if *service == "" {
		fmt.Println("You need to define a service to monitor.")
	} else {
		serviceString := string(*service)
		exporter := NewExporter(serviceString)
		prometheus.MustRegister(exporter)
	}

	log.Printf("Starting Server: %s\n", s.ListenAddress)
	log.Printf("Metrics Path: %s\n", s.MetricsPath)
	handler := prometheus.Handler()

	if s.MetricsPath == "" || s.MetricsPath == "/" {
		http.Handle(s.MetricsPath, handler)
	} else {
		http.Handle(s.MetricsPath, handler)
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(`<html>
				<head><title>Prometheus Service Exporter</title></head>
				<body>
				<h1>Prometheus Service Exporter</h1>
				<p><a href="` + s.MetricsPath + `">Metrics</a></p>
				</body>
				</html>`))
		})
	}
	err = http.ListenAndServe(s.ListenAddress, nil)
	if err != nil {
		log.Fatal(err)
	}
}
