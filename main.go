package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"math/rand"
	"net/http"
)

var onlineUsers = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "app_go_online_users",
	Help: "Online users",
	ConstLabels: map[string]string{
		"course": "fullcyle",
	},
})

func main() {
	r := prometheus.NewRegistry()
	r.MustRegister(onlineUsers)

	go func() {
		for {
			onlineUsers.Set(float64(rand.Intn(700)))
		}
	}()

	http.Handle("/metrics", promhttp.HandlerFor(r, promhttp.HandlerOpts{}))
	log.Fatal(http.ListenAndServe(":8181", nil))
}
