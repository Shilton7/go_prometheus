package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"math/rand"
	"net/http"
	"time"
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

// Histogram
var httpDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Name: "app_go_http_request_duration",
	Help: "Duration in seconds of all HTTP requests for appgo",
}, []string{"handler"})

func main() {
	r := prometheus.NewRegistry()

	r.MustRegister(onlineUsers)       // Gauge
	r.MustRegister(httpRequestsTotal) // Counter
	r.MustRegister(httpDuration)      // Histogram

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
		time.Sleep(time.Duration(rand.Intn(7)) * time.Second)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello Shilton!"))
	})

	// Counter old
	// http.Handle("/", promhttp.InstrumentHandlerCounter(httpRequestsTotal, home))

	// Histogram
	durationHome := promhttp.InstrumentHandlerDuration(
		httpDuration.MustCurryWith(prometheus.Labels{"handler": "home"}),
		promhttp.InstrumentHandlerCounter(httpRequestsTotal, home),
	)

	// Histogram + Counter
	http.Handle("/", durationHome)

	// Contact
	contact := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Duration(rand.Intn(12)) * time.Second)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Contact Here!"))
	})

	durationContact := promhttp.InstrumentHandlerDuration(
		httpDuration.MustCurryWith(prometheus.Labels{"handler": "contact"}),
		promhttp.InstrumentHandlerCounter(httpRequestsTotal, contact),
	)

	http.Handle("/contact", durationContact)

	// Server Start
	log.Fatal(http.ListenAndServe(":8181", nil))
}
