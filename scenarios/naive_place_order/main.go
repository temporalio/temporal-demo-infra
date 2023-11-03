package main

import (
	"context"
	"github.com/temporalio/temporal-demo-infra/internal/clients"
	temporal2 "github.com/temporalio/temporal-demo-infra/pkg/clients/temporal"
	"github.com/temporalio/temporal-demo-infra/pkg/config"
	"github.com/temporalio/temporal-demo-infra/pkg/instrumentation/log"
	"logur.dev/logur"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type startable interface {
	Start(context.Context) error
	Shutdown(context.Context)
}

type appConfig struct {
	Log            *log.Config
	TemporalClient *temporal2.Config
	Scenario       *Config
}

func main() {
	// config root
	config.MustLoad()
	var err error

	appCfg := &appConfig{}
	config.MustUnmarshalAll(appCfg)

	ctx, _ := context.WithCancel(context.Background())

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

	// clients
	clients := clients.MustGetClients(ctx,
		clients.WithTemporal(temporal2.NewClients(ctx,
			temporal2.WithConfig(appCfg.TemporalClient),
			temporal2.WithLogger(logger))),
	)

	defer func() {
		if perr := clients.Close(); perr != nil {
			logger.Error("failed to close clients gracefully", logur.Fields{"err": perr})
		}
	}()
	if err := performScenario(ctx, appCfg, clients); err != nil {
		panic("scenario failed : " + err.Error())
	}
}

type Config struct {
	TaskQueueOrders string
}

func (c *Config) Prefix() string {
	return "scenario"
}
func wait(secs int64) {
	select {
	case <-time.After(time.Duration(secs) * time.Second):
		return
	}
}
