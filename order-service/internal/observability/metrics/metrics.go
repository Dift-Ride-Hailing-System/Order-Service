package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Metrics holds all the metrics for the service
type Metrics struct {
	OrdersCreated    prometheus.Counter
	OrdersCancelled  prometheus.Counter
	OrdersConfirmed  prometheus.Counter
	OrderProcessTime prometheus.Histogram
}

var m *Metrics

// InitMetrics initialize all metrics
func InitMetrics() *Metrics {
	if m != nil {
		return m
	}

	m = &Metrics{
		OrdersCreated: promauto.NewCounter(prometheus.CounterOpts{
			Name: "order_created_total",
			Help: "Total number of orders created",
		}),
		OrdersCancelled: promauto.NewCounter(prometheus.CounterOpts{
			Name: "order_cancelled_total",
			Help: "Total number of orders cancelled",
		}),
		OrdersConfirmed: promauto.NewCounter(prometheus.CounterOpts{
			Name: "order_confirmed_total",
			Help: "Total number of orders confirmed",
		}),
		OrderProcessTime: promauto.NewHistogram(prometheus.HistogramOpts{
			Name:    "order_process_duration_seconds",
			Help:    "Time taken to process orders",
			Buckets: prometheus.DefBuckets,
		}),
	}

	return m
}
