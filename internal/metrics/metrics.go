package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	OpsProcessed = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "reconcile_total_number_loops_counter",
		Help: "The total number of processed events",
	})

	OpsFailed = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "reconcile_total_number_failed_loops_counter",
		Help: "How many errors happened during run time of loop",
	})
)

func init() {
	prometheus.MustRegister(OpsProcessed)
	prometheus.MustRegister(OpsFailed)
}
