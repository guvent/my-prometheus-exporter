package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"time"
)

type fooCollector struct {
	fooMetric *prometheus.Desc
	barMetric *prometheus.Desc
}

var metricValue float64

func newFooCollector() *fooCollector {
	return &fooCollector{
		fooMetric: prometheus.NewDesc("foo_metric",
			"Shows whether a foo has occurred in our cluster 1",
			nil, nil,
		),
		barMetric: prometheus.NewDesc("bar_metric",
			"Shows whether a bar has occurred in our cluster 2",
			nil, nil,
		),
	}
}

func (collector *fooCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.fooMetric
	ch <- collector.barMetric
}

func (collector *fooCollector) Collect(ch chan<- prometheus.Metric) {
	m1 := prometheus.MustNewConstMetric(collector.fooMetric, prometheus.GaugeValue, metricValue)
	m2 := prometheus.MustNewConstMetric(collector.barMetric, prometheus.GaugeValue, metricValue)
	m1 = prometheus.NewMetricWithTimestamp(time.Now().Add(-time.Hour), m1)
	m2 = prometheus.NewMetricWithTimestamp(time.Now(), m2)
	ch <- m1
	ch <- m2
}


func main() {
	foo := newFooCollector()
	prometheus.MustRegister(foo)

	http.HandleFunc("/console/set", func(w http.ResponseWriter, r *http.Request) {
		metricValue = 10

		w.Header().Add("Location", "/")
		w.WriteHeader(302)

		if _, err := w.Write([]byte(``)); err != nil { log.Fatal(err) }
	})

	http.HandleFunc("/console/pick", func(w http.ResponseWriter, r *http.Request) {
		metricValue = 10
		time.Sleep(time.Second*1)
		metricValue = 0

		w.Header().Add("Location", "/")
		w.WriteHeader(302)

		if _, err := w.Write([]byte(``)); err != nil { log.Fatal(err) }
	})

	http.HandleFunc("/console/reset", func(w http.ResponseWriter, r *http.Request) {
		metricValue = 0

		w.Header().Add("Location", "/")
		w.WriteHeader(302)

		if _, err := w.Write([]byte(``)); err != nil { log.Fatal(err) }
	})

	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte(`
            <html>
            <head><title>Volume Exporter Metrics</title></head>
            <body>
            <h1>ConfigMap Reload</h1>
            <p><a href='/metrics'>Show metrics data</a></p>
            <p><a href='/console/set'>Set value is 10</a></p>
            <p><a href='/console/pick'>Pick value 1 sec. is 10</a></p>
            <p><a href='/console/reset'>Reset value to 0</a></p>
            </body>
            </html>
        `)); err != nil {
			return
		}
	})

	log.Fatal(http.ListenAndServe(":9888", nil))
}
