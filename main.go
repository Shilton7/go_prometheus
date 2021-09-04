package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"math/rand"
	"net/http"
)

// Gauge
var onlineUsers = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "app_go_online_users",
	Help: "Online users",
	ConstLabels: map[string]string{
		"course": "fullcyle",
	},
})

// Counter
var httpRequestsTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "app_go_http_requests_total",
	Help: "Count of all HTTP requests for appgo",
}, []string{})

func main() {
	r := prometheus.NewRegistry()

	r.MustRegister(onlineUsers)       //Gauge
	r.MustRegister(httpRequestsTotal) // Counter

	//Gauge
	go func() {
		for {
			onlineUsers.Set(float64(rand.Intn(700)))
		}
	}()

	// Gauge
	http.Handle("/metrics", promhttp.HandlerFor(r, promhttp.HandlerOpts{}))

	// Counter
	home := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello Shilton!"))
	})

	// Counter
	http.Handle("/", promhttp.InstrumentHandlerCounter(httpRequestsTotal, home))

	// Server Start
	log.Fatal(http.ListenAndServe(":8181", nil))
}
