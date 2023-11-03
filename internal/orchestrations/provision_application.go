package orchestrations

import (
	"errors"
	"fmt"
	"github.com/mrz1836/go-sanitize"
	"github.com/temporalio/temporal-demo-infra/internal/notifications"
	"github.com/temporalio/temporal-demo-infra/internal/teams"
	"github.com/temporalio/temporal-demo-infra/pkg/messages"
	"go.temporal.io/sdk/log"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
	"strings"
	"time"
)

type CloudProvisionerInfo struct {
	TaskQueue         string
	ProvisionActivity string
	DestroyActivity   string
}

var AWSProvisioner = CloudProvisionerInfo{
	TaskQueue:         "provisioning_aws",
	ProvisionActivity: "provisionFoundationResources",
	DestroyActivity:   "destroyFoundationResources",
}

type ProvisionApplicationState struct {
	AuthorizationTimedOut bool
	Authorization         *messages.AuthorizationReceivedResponse
	Resources             *messages.ProvisionFoundationResourcesResponse
}

func (s *ProvisionApplicationState) ToKVP() []interface{} {
	return []interface{}{
		"authorization_approved", s.Authorization != nil && s.Authorization.IsApproved,
		"resources_created", s.Resources != nil,
		"authorization_timed_out", s.AuthorizationTimedOut,
	}
}

func (o *Orchestrations) ProvisionApplication(ctx workflow.Context, params *messages.ProvisionApplicationRequest) (*ProvisionApplicationState, error) {

	if params.AuthorizationTimeoutSeconds == 0 {
		params.AuthorizationTimeoutSeconds = 120
	}
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		// activities should be completed within this time
		StartToCloseTimeout: time.Second * 3,
	})
	logger := log.With(
		workflow.GetLogger(ctx),
		"teamId",
		params.TeamID,
		"applicationId",
		params.ApplicationID,
		"applicationName",
		params.ApplicationName,
	)
	logger.Info("provisioning application")

	// get details and execute activities to do the work
	var teamInfo *messages.GetTeamInformationResponse

	state := &ProvisionApplicationState{}

	authorizationCtx, _ := workflow.WithCancel(ctx)
	provisionCtx, _ := workflow.WithCancel(workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		// 5 minutes to provision our resources
		StartToCloseTimeout: time.Second * 3600,
		// ping Temporal every 10 seconds that we are still working
		HeartbeatTimeout: time.Second * 10,
		// special sauce for targeting polyglot
		TaskQueue: AWSProvisioner.TaskQueue,
	}))
	defer (func() {
		if !errors.Is(ctx.Err(), workflow.ErrCanceled) {
			return
		}
		// no resources were provisioned so just exit
		if state.Resources == nil {
			return
		}
		// When the Workflow is canceled, it has to get a new disconnected context to execute any Activities
		newCtx, _ := workflow.NewDisconnectedContext(ctx)
		err := workflow.ExecuteActivity(newCtx, "destroyResources", state.Resources).Get(ctx, nil)
		if err != nil {
			logger.Error("CleanupActivity failed", "Error", err)
		}
	})()

	// setup listeners for authorization
	authorizationChan := workflow.GetSignalChannel(ctx, messages.MessageName(state.Authorization))
	authHandlerCtx, done := workflow.WithCancel(ctx)
	workflow.GoNamed(authHandlerCtx, "authorize", func(ctx workflow.Context) {
		logger.Info("listening for authorization and fulfillment events")

		for {
			s := workflow.NewNamedSelector(ctx, "authorize").
				AddReceive(authorizationChan, func(c workflow.ReceiveChannel, more bool) {
					c.Receive(ctx, &state.Authorization)
					logger.Info("authorization received", "is_approved", state.Authorization.IsApproved)

				})
			s.Select(ctx)
			if state.Authorization != nil && state.Resources != nil {
				done()
			}
		}
	})

	if err := workflow.ExecuteActivity(
		ctx,
		teams.TypeHandlers.GetTeamInformation,
		&messages.GetTeamInformationRequest{
			RequesterID: params.RequesterID,
			TeamID:      params.TeamID,
		},
	).Get(ctx, &teamInfo); err != nil {
		return nil, fmt.Errorf("failed to get team information %w", err)
	}

	logger.Info("received provision request", "request", teamInfo)

	// this is a fire and forget call, so we can block for the successful send of an API request
	if err := workflow.ExecuteActivity(
		authorizationCtx,
		notifications.TypeHandlers.RequestApplicationAuthorization,
		&messages.RequestApplicationAuthorizationRequest{
			AuthorizerID:  params.AuthorizerID,
			TeamID:        params.TeamID,
			CustomerID:    params.RequesterID,
			DelaySeconds:  params.DemoAuthorizationDelaySeconds,
			ApplicationID: params.ApplicationID,
		},
	).Get(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to send authorization request %w", err)
	}

	// BLOCK for all our conditions to be met before proceeding
	logger.Info("waiting for authorization and fulfillment to complete")
	_, err := workflow.AwaitWithTimeout(ctx, time.Second*time.Duration(params.AuthorizationTimeoutSeconds), func() bool {
		// allow time for approval to be changed so only block while not approved
		return state.Authorization != nil
	})

	state.AuthorizationTimedOut = temporal.IsCanceledError(err)
	logger.Info("done waiting", state.ToKVP()...)
	if state.Authorization == nil || !state.Authorization.IsApproved {
		// no soup for you
		return nil, fmt.Errorf("application '%s' is not approved", params.ApplicationName)
	}

	if state.Authorization != nil {
		// this is a fire and forget call, so we can block for the successful send of an API request
		if err := workflow.ExecuteActivity(
			provisionCtx,
			AWSProvisioner.ProvisionActivity,
			&messages.ProvisionFoundationResourcesRequest{
				ApplicationID:   params.ApplicationID,
				TeamID:          params.TeamID,
				Region:          state.Authorization.Region,
				Profile:         state.Authorization.Profile,
				ApplicationName: params.ApplicationName,
				BucketName: fmt.Sprintf(
					"app-%s-%s-%s",
					strings.ToLower(sanitize.AlphaNumeric(params.ApplicationName, false)),
					strings.ToLower(sanitize.AlphaNumeric(params.TeamID, false)),
					strings.ToLower(sanitize.AlphaNumeric(params.ApplicationID, false)),
				),
			},
		).Get(ctx, &state.Resources); err != nil {
			return nil, fmt.Errorf("failed to provision %w", err)
		}
	} else if temporal.IsCanceledError(err) {
		return nil, fmt.Errorf("authorization was not achieved in time")
	}

	logger.Info("application has been completed and authorized", "application completed on", state.Resources.CompletionDateTime)
	drain(authorizationChan, state.Authorization)
	return state, nil
}
func drain(signalChan workflow.ReceiveChannel, obj interface{}) {
	// Drain signal channel asynchronously to avoid signal loss
	for {
		ok := signalChan.ReceiveAsync(&obj)
		if !ok {
			break
		}
	}
}
