package notifications

import (
	"context"
	"fmt"
	"github.com/temporalio/temporal-demo-infra/pkg/messages"
	"go.temporal.io/sdk/activity"
	sdkclient "go.temporal.io/sdk/client"
	"time"
)

var TypeHandlers *Handlers

type Handlers struct {
	temporalClient sdkclient.Client
}

func NewHandlers(temporalClient sdkclient.Client) *Handlers {
	return &Handlers{temporalClient: temporalClient}
}

// RequestApplicationAuthorization sends a message (eg email) to the Authorizer (eg Doctor) to obtain approval for fulfillment
func (h *Handlers) RequestApplicationAuthorization(ctx context.Context, cmd *messages.RequestApplicationAuthorizationRequest) error {
	logger := activity.GetLogger(ctx)
	logger.Info(
		"authorizing application",
		"applicationId",
		cmd.ApplicationID,
		"teamId",
		cmd.TeamID,
	)

	if cmd.DelaySeconds == 0 {
		return nil
	}
	go (func(cmd *messages.RequestApplicationAuthorizationRequest) {
		select {
		case <-time.After(time.Duration(cmd.DelaySeconds) * time.Second):
			event := &messages.AuthorizationReceivedResponse{
				ApplicationID: cmd.ApplicationID,
				Region:        "us-east-1",
				Profile:       "iac",
				TeamID:        cmd.TeamID,
				IsApproved:    true,
			}
			if err := h.temporalClient.SignalWorkflow(context.Background(), cmd.ApplicationID, "", messages.MessageName(event), event); err != nil {
				fmt.Println(fmt.Errorf("failed to authorize prescription %w", err))
			}
		}
	})(cmd)
	return nil
}
