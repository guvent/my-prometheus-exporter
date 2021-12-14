package utils

import "github.com/prometheus/client_golang/prometheus"

type volumeCollector struct {
	volumeBytesTotal *prometheus.Desc
	// More....
}

func newVolumeCollector(namespace string) *volumeCollector {
	return &volumeCollector{
		volumeBytesTotal: prometheus.NewDesc(prometheus.BuildFQName(namespace, "", "bytes_total"),
			"Total size of the volume/disk",
			[]string{"name", "path"}, nil,
		),
		// More....
	}
}

func (collector *volumeCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.volumeBytesTotal
	// More....
}

func (collector *volumeCollector) Collect(ch chan<- prometheus.Metric) {
	var metricValue float64

	if 1 == 1 { metricValue = 1 }

	ch <- prometheus.MustNewConstMetric(collector.volumeBytesTotal, prometheus.GaugeValue, metricValue, "log", "path")
	// More....
}

func Register() {
	collector := newVolumeCollector("my_metric")
	prometheus.MustRegister(collector)
}