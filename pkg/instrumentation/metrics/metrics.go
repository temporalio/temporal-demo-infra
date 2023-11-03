package metrics

import (
	"context"
	"fmt"
	prom "github.com/prometheus/client_golang/prometheus"
	log2 "github.com/temporalio/temporal-demo-infra/pkg/instrumentation/log"
	"github.com/uber-go/tally/v4"
	"github.com/uber-go/tally/v4/prometheus"
	sdktally "go.temporal.io/sdk/contrib/tally"
	"time"
)

func NewPrometheusScope(ctx context.Context, cfg *Config) (tally.Scope, error) {

	if cfg.TimerType == "" {
		cfg.TimerType = "histogram"
	}
	logger := log2.WithFields(log2.GetLogger(ctx), log2.Fields{
		"listen_address": cfg.ListenAddress,
		"timer_type":     cfg.TimerType,
		"prefix":         cfg.ScopePrefix,
	})
	boundaries := []float64{.001, .005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10, 25, 60, 70}
	buckets := make([]prometheus.HistogramObjective, 0, len(boundaries))
	for _, boundary := range boundaries {
		buckets = append(buckets, prometheus.HistogramObjective{Upper: boundary})
	}
	c := prometheus.Configuration{
		ListenAddress:           cfg.ListenAddress,
		TimerType:               cfg.TimerType,
		DefaultHistogramBuckets: buckets,
	}
	reporter, err := c.NewReporter(
		prometheus.ConfigurationOptions{
			Registry: prom.NewRegistry(),
			OnError: func(err error) {
				logger.Error("error in prometheus reporter", log2.Fields{"err": err})
			},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("error creating prometheus reporter %w", err)
	}
	scopeOpts := tally.ScopeOptions{
		CachedReporter:  reporter,
		Separator:       prometheus.DefaultSeparator,
		SanitizeOptions: &sdktally.PrometheusSanitizeOptions,
		Prefix:          cfg.ScopePrefix,
	}
	scope, _ := tally.NewRootScope(scopeOpts, time.Second)
	scope = sdktally.NewPrometheusNamingScope(scope)
	logger.Info("prometheus metrics scope created")
	return scope, nil
}
