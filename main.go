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

	http.HandleFunc("/console/up", func(w http.ResponseWriter, r *http.Request) {
		metricValue++

		w.Header().Add("Location", "/")
		w.WriteHeader(302)

		if _, err := w.Write([]byte(``)); err != nil {
			log.Fatal(err)
		}
	})

	http.HandleFunc("/console/down", func(w http.ResponseWriter, r *http.Request) {
		metricValue--

		w.Header().Add("Location", "/")
		w.WriteHeader(302)

		if _, err := w.Write([]byte(``)); err != nil {
			log.Fatal(err)
		}
	})

	http.HandleFunc("/console/up2", func(w http.ResponseWriter, r *http.Request) {
		metricValue = +10

		w.Header().Add("Location", "/")
		w.WriteHeader(302)

		if _, err := w.Write([]byte(``)); err != nil {
			log.Fatal(err)
		}
	})

	http.HandleFunc("/console/down2", func(w http.ResponseWriter, r *http.Request) {
		metricValue = -10

		w.Header().Add("Location", "/")
		w.WriteHeader(302)

		if _, err := w.Write([]byte(``)); err != nil {
			log.Fatal(err)
		}
	})

	http.HandleFunc("/console/pick", func(w http.ResponseWriter, r *http.Request) {
		metricValue = 20
		time.Sleep(time.Second * 20)
		metricValue = 0

		w.Header().Add("Location", "/")
		w.WriteHeader(302)

		if _, err := w.Write([]byte(``)); err != nil {
			log.Fatal(err)
		}
	})

	http.HandleFunc("/console/reset", func(w http.ResponseWriter, r *http.Request) {
		metricValue = 0

		w.Header().Add("Location", "/")
		w.WriteHeader(302)

		if _, err := w.Write([]byte(``)); err != nil {
			log.Fatal(err)
		}
	})

	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte(`
            <html>
            <head><title>My Exporter Metrics</title></head>
            <body>
            <h1>ConfigMap Reload</h1>
            <p><a href='/metrics'>Show metrics data</a></p>
            <p><a href='/console/up'>Value increment 1</a></p>
            <p><a href='/console/down'>Value decrement 1</a></p>
            <p><a href='/console/up2'>Value increment 10</a></p>
            <p><a href='/console/down2'>Value decrement 10</a></p>
            <p><a href='/console/pick'>Pick value 20 sec. as 10</a></p>
            <p><a href='/console/reset'>Reset value to 0</a></p>
            </body>
            </html>
        `)); err != nil {
			return
		}
	})

	log.Fatal(http.ListenAndServe(":9888", nil))
}
