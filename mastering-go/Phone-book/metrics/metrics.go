package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var METRICS_PORT = ":1234"

func RegisterMetrics() []prometheus.Collector {
	counterMetric := prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "phone_book",
			Name:      "counter_metric",
		},
	)

	gaugeMetric := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "phone_book",
			Name:      "gauge_metric",
		},
	)

	histogramMetric := prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: "phone_book",
			Name:      "histogram_metric",
		},
	)

	summaryMetric := prometheus.NewSummary(
		prometheus.SummaryOpts{
			Namespace: "phone_book",
			Name:      "summary_metric",
		},
	)

	counterMetric.Add(5)

	return []prometheus.Collector{summaryMetric, histogramMetric, gaugeMetric, counterMetric}
}
