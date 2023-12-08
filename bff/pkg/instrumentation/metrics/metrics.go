package metrics

import (
	"context"
	"fmt"
	prom "github.com/prometheus/client_golang/prometheus"
	log2 "github.com/temporalio/temporal-demo-infra/bff/pkg/instrumentation/log"
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
	c := prometheus.Configuration{
		ListenAddress: cfg.ListenAddress,
		TimerType:     cfg.TimerType,
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
