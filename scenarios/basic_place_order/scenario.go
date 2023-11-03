package main

import (
	"context"
	"fmt"
	"github.com/temporalio/temporal-demo-infra/internal/clients"
	"github.com/temporalio/temporal-demo-infra/internal/orchestrations"
	"github.com/temporalio/temporal-demo-infra/pkg/instrumentation/log"
	"github.com/temporalio/temporal-demo-infra/pkg/messages"
	"github.com/teris-io/shortid"
	"go.temporal.io/sdk/client"
)

func performScenario(ctx context.Context, cfg *appConfig, clients *clients.Clients) error {
	workflowID := shortid.MustGenerate()
	logger := log.WithFields(log.GetLogger(ctx), log.Fields{"workflow_id": workflowID})
	var err error
	var run client.WorkflowRun

	params := &messages.ProvisionApplicationRequest{
		ApplicationID:                 workflowID,
		RequesterID:                   shortid.MustGenerate(),
		TeamID:                        shortid.MustGenerate(),
		ApplicationName:               shortid.MustGenerate(),
		AuthorizerID:                  shortid.MustGenerate(),
		AuthorizationTimeoutSeconds:   6,
		DemoAuthorizationDelaySeconds: 2,
		DemoFulfillmentDelaySeconds:   3,
	}
	logger.Info("starting scenario")
	temporalClient := clients.Temporal().Client
	run, err = temporalClient.ExecuteWorkflow(ctx, client.StartWorkflowOptions{
		ID:        workflowID,
		TaskQueue: cfg.Scenario.TaskQueueOrders,
	}, orchestrations.TypeOrchestrations.ProvisionApplication, params)
	logger.Info("started scenario", log.Fields{"run_id": run.GetRunID()})
	if err != nil {
		return err
	}

	logger.Info("waiting for completion of order")
	// block for completion of order
	if err = run.Get(ctx, nil); err != nil {
		return fmt.Errorf("order error %w", err)
	}
	logger.Info("order completed!")
	return nil
}
