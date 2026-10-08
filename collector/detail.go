package collector

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/percona/mongodb_exporter/collector/mongod"
	"github.com/percona/mongodb_exporter/collector/serverstatusv5"
)

// NewDetailedHandler returns a modern server-status handler with persistent
// diagnostic counters and request-local raw response handling.
func NewDetailedHandler(exporter *MongodbCollector) http.Handler {
	return newDetailedHandler(
		func() *mongo.Client { return exporter.getClient() },
		mongod.GetRawServerStatus,
		exporter.Opts.SuppressServerStatusV5,
	)
}

type detailedClientGetter func() *mongo.Client
type detailedRawGetter func(*mongo.Client) (bson.Raw, error)

func newDetailedHandler(getClient detailedClientGetter, getRaw detailedRawGetter, suppressed bool) http.Handler {
	module := serverstatusv5.NewDetailed()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if suppressed {
			http.Error(w, "detailed MongoDB server-status collection is disabled", http.StatusServiceUnavailable)
			return
		}

		client := getClient()
		if client == nil {
			http.Error(w, "unable to create MongoDB client", http.StatusServiceUnavailable)
			return
		}

		raw, err := getRaw(client)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to collect MongoDB server status: %v", err), http.StatusInternalServerError)
			return
		}

		if process, ok := raw.Lookup("process").StringValueOK(); ok && process == "mongos" {
			http.Error(w, "detailed MongoDB server-status collection is unavailable on mongos", http.StatusNotImplemented)
			return
		}

		newDetailedMetricsHandler(module, raw).ServeHTTP(w, r)
	})
}

// newDetailedMetricsHandler creates a fresh modern-only registry and raw
// collector for one response. The handler supplies its reusable module so
// diagnostics persist, while the registry and copied raw stay request-local.
func newDetailedMetricsHandler(module *serverstatusv5.Module, raw bson.Raw) http.Handler {
	registry := prometheus.NewRegistry()
	collector := &detailedRawCollector{
		module: module,
		raw:    append(bson.Raw(nil), raw...),
	}
	if err := registry.Register(collector); err != nil {
		return http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			http.Error(w, fmt.Sprintf("failed to register detailed metrics: %v", err), http.StatusInternalServerError)
		})
	}
	return promhttp.HandlerFor(registry, promhttp.HandlerOpts{ErrorHandling: promhttp.HTTPErrorOnError})
}

type detailedRawCollector struct {
	module *serverstatusv5.Module
	raw    bson.Raw
}

func (collector *detailedRawCollector) Describe(ch chan<- *prometheus.Desc) {
	collector.module.Describe(ch)
}

func (collector *detailedRawCollector) Collect(ch chan<- prometheus.Metric) {
	collector.module.Collect(collector.raw, ch)
}

var _ prometheus.Collector = (*detailedRawCollector)(nil)
