package temporal_demo_infra

import (
	"context"
	"fmt"
	"github.com/temporalio/temporal-demo-infra/internal/notifications"
	"github.com/temporalio/temporal-demo-infra/internal/orchestrations"
	"github.com/temporalio/temporal-demo-infra/internal/teams"
	temporalClient "github.com/temporalio/temporal-demo-infra/pkg/clients/temporal"
	"time"

	"github.com/temporalio/temporal-demo-infra/pkg/instrumentation/log"

	"logur.dev/logur"

	"go.temporal.io/sdk/worker"
)

const TaskQueueApps = "apps"

type Option func(w *Worker)

func WithTemporal(c *temporalClient.Clients) Option {
	return func(w *Worker) {
		w.temporalClients = c
	}
}
func WithConfig(cfg *Config) Option {
	return func(w *Worker) {
		w.cfg = cfg
	}
}

type Config struct {
	TaskQueueApps string
}

func (c *Config) Prefix() string {
	return "worker"
}

type Worker struct {
	temporalClients *temporalClient.Clients
	inner           worker.Worker
	cfg             *Config

	// other clients
}

func NewWorker(_ context.Context, opts ...Option) (*Worker, error) {
	w := &Worker{}
	defaultOpts := []Option{WithConfig(&Config{TaskQueueApps: TaskQueueApps})}
	opts = append(defaultOpts, opts...)
	for _, o := range opts {
		o(w)
	}
	return w, nil
}
func (w *Worker) Shutdown(_ context.Context) {
	// TODO
}

// register wires up dependencies and registers activities and/or workflows onto underlying worker
func (w *Worker) register(inner worker.Worker) error {

	wfs := &orchestrations.Orchestrations{}

	// register activities
	teamsHandlers := teams.NewHandlers()
	notificationsHandlers := notifications.NewHandlers(w.temporalClients.Client)

	inner.RegisterActivity(teamsHandlers)
	inner.RegisterActivity(notificationsHandlers)

	// register workflows
	inner.RegisterWorkflow(wfs.ProvisionApplication)
	return nil
}
func (w *Worker) Start(ctx context.Context) error {
	logger := log.GetLogger(ctx)

	inner := worker.New(w.temporalClients.Client, w.cfg.TaskQueueApps, worker.Options{})

	if err := w.register(inner); err != nil {
		return fmt.Errorf("failed to register workflows/activities: %w", err)
	}

	w.inner = inner
	err := w.inner.Run(worker.InterruptCh())
	if err != nil {
		logger.Error("start worker failed", logur.Fields{"err": err})
		return err
	}
	return nil
}

// https://legacy-documentation-sdks.temporal.io/go/how-to-set-workeroptions-in-go
// https://pkg.go.dev/go.temporal.io/sdk@v1.23.0/internal#WorkerOptions
func configureWorker() worker.Options {
	// cache
	worker.SetStickyWorkflowCacheSize(10000)
	// defaults are assigned here
	workerOptions := worker.Options{
		// Pollers
		MaxConcurrentActivityTaskPollers: 2,
		MaxConcurrentWorkflowTaskPollers: 2,

		// Executors
		MaxConcurrentActivityExecutionSize:      1000,
		LocalActivityWorkerOnly:                 false,
		MaxConcurrentLocalActivityExecutionSize: 1000,
		MaxConcurrentEagerActivityExecutionSize: 0, //unlimited
		MaxConcurrentWorkflowTaskExecutionSize:  1000,
		MaxConcurrentSessionExecutionSize:       1000,

		// Cache
		StickyScheduleToStartTimeout: time.Second * 5,

		// Heartbeat
		MaxHeartbeatThrottleInterval:     60, // seconds
		DefaultHeartbeatThrottleInterval: 30, // seconds

		//Rate Limiting
		WorkerActivitiesPerSecond:      10000,
		WorkerLocalActivitiesPerSecond: 10000,
		TaskQueueActivitiesPerSecond:   10000,

		// Worker Type
		EnableSessionWorker:   false,
		DisableWorkflowWorker: false,

		// Worker ID
		Identity: "",

		// Versioning
		BuildID:                 "",
		UseBuildIDForVersioning: false,

		// Runtime Behavior
		EnableLoggingInReplay:       false,
		BackgroundActivityContext:   nil,
		WorkflowPanicPolicy:         worker.BlockWorkflow,
		WorkerStopTimeout:           0,
		DeadlockDetectionTimeout:    time.Second * 1,
		Interceptors:                nil,
		OnFatalError:                nil,
		DisableRegistrationAliasing: false, // recommended
		DisableEagerActivities:      false,
	}
	return workerOptions
}
