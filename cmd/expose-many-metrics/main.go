package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	numTimeseries = flag.Int("timeseries", 10000, "The number of timeseries to return.")
	metric        = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "example_timeseries_total",
		Help: "Timeseries being exposed",
	},
		[]string{"long_label"},
	)
)

type collector struct {
}

func randFloat() float64 {
	// Simulate nanosecond precision latency around 1s.
	r := rand.NormFloat64() + 1
	r *= 1e9
	return math.Round(r) / 1e9
}

func (collector) Describe(ch chan<- *prometheus.Desc) {
	metric.Describe(ch)
}

func (collector) Collect(ch chan<- prometheus.Metric) {
	for i := 0; i < *numTimeseries; i++ {
		metric.WithLabelValues(
			fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s%d",
				"abcdefghijklmnopqrstuvwxyz",
				"abcdefghijklmnopqrstuvwxyz",
				"abcdefghijklmnopqrstuvwxyz",
				"abcdefghijklmnopqrstuvwxyz",
				"abcdefghijklmnopqrstuvwxyz",
				"abcdefghijklmnopqrstuvwxyz",
				"abcdefghijklmnopqrstuvwxyz",
				"abcdefghijklmnopqrstuvwxyz",
				"abcdefghijklmnopqrstuvwxyz",
				"abcdefghijklmnopqrstuvwxyz", i),
		).Add(math.Max(0, randFloat()))
	}
	metric.Collect(ch)
}

func main() {
	prometheus.MustRegister(collector{})
	flag.Parse()
	log.Printf("Quickly generating %d timeseries....", *numTimeseries)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":12345", nil)
}
