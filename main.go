package main

import (
	"context"
	"fmt"
	temporal_demo_infra "github.com/temporalio/temporal-demo-infra/internal/workers/temporal"
	temporal2 "github.com/temporalio/temporal-demo-infra/pkg/clients/temporal"
	"github.com/temporalio/temporal-demo-infra/pkg/instrumentation/metrics"
	"github.com/temporalio/temporal-demo-infra/pkg/instrumentation/probes"
	sdkclient "go.temporal.io/sdk/client"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/temporalio/temporal-demo-infra/internal/clients"
	"github.com/temporalio/temporal-demo-infra/pkg/config"
	"github.com/temporalio/temporal-demo-infra/pkg/instrumentation/log"
	"golang.org/x/sync/errgroup"
	"logur.dev/logur"
)

type startable interface {
	Start(context.Context) error
	Shutdown(context.Context)
}

type appConfig struct {
	Log            *log.Config
	TemporalClient *temporal2.Config
	TemporalWorker *temporal_demo_infra.Config
	Metrics        *metrics.Config
}

func main() {
	// config root
	config.MustLoad()
	var err error

	appCfg := &appConfig{}
	config.MustUnmarshalAll(appCfg)

	ctx, done := context.WithCancel(context.Background())
	g, ctx := errgroup.WithContext(ctx)

	// set up signal listener
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(quit)

	// logging
	var logger logur.Logger
	if logger, err = log.NewLogger(ctx, appCfg.Log); err != nil {
		panic("failed to create logger" + err.Error())
	}
	ctx = log.WithLogger(ctx, logger)
	// apps

	mts, err := metrics.NewPrometheusScope(ctx, appCfg.Metrics)
	if err != nil {
		panic(fmt.Sprintf("failed to create prometheus scope %s", err.Error()))
	}
	logger.Info("prometheus scope created", log.Fields{"scope": mts})
	// clients
	clients := clients.MustGetClients(ctx,
		clients.WithTemporal(temporal2.NewClients(ctx,
			temporal2.WithConfig(appCfg.TemporalClient),
			temporal2.WithOptions(sdkclient.Options{
				MetricsHandler: temporal2.NewMetricsHandler(mts),
			}),
			temporal2.WithLogger(logger))),
	)

	defer func() {
		if perr := clients.Close(); perr != nil {
			logger.Error("failed to close clients gracefully", logur.Fields{"err": perr})
		}
	}()

	wk, err := temporal_demo_infra.NewWorker(
		ctx,
		temporal_demo_infra.WithTemporal(clients.Temporal()),
		temporal_demo_infra.WithConfig(appCfg.TemporalWorker),
	)
	if err != nil {
		panic(fmt.Errorf("failed to create temporal worker: %v", err))
	}

	probeServer, err := probes.NewProbes(ctx)
	if err != nil {
		panic("failed to create probes: " + err.Error())
	}

	startables := []startable{probeServer, wk}

	for _, s := range startables {
		var current = s
		g.Go(func() error {
			if err := current.Start(ctx); err != nil {
				return err
			}
			return nil
		})
	}

	select {
	case <-quit:
		break
	case <-ctx.Done():
		break
	}

	// shutdown the things
	done()

	// limit how long we'll wait for
	timeoutCtx, timeoutCancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer timeoutCancel()

	logger.Info("shutting down servers, please wait...")

	for _, s := range startables {
		s.Shutdown(timeoutCtx)
	}

	// wait for shutdown
	if err := g.Wait(); err != nil {
		panic("shutdown was not clean" + err.Error())
	}
	logger.Info("goodbye")
}
