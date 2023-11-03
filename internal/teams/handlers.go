package teams

import (
	"context"
	"fmt"
	"github.com/temporalio/temporal-demo-infra/pkg/messages"
	"github.com/teris-io/shortid"
	"go.temporal.io/sdk/activity"
	"time"
)

var TypeHandlers *Handlers

type Handlers struct {
}

func NewHandlers() *Handlers {
	return &Handlers{}
}

// GetTeamInformation simulates a lookup (eg database reads) that DO NOT MUTATE application state
func (h *Handlers) GetTeamInformation(ctx context.Context, request *messages.GetTeamInformationRequest) (*messages.GetTeamInformationResponse, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("GetTeamInformation", "request", request)
	return &messages.GetTeamInformationResponse{
		TeamName:  shortid.MustGenerate(),
		TeamEmail: fmt.Sprintf("team-%s@example.com", request.TeamID),
		Subdomain: "apps_marketing",
	}, nil
}
func (h *Handlers) ValidateInsurance(ctx context.Context, cmd *messages.ValidateInsuranceRequest) (*messages.ValidateInsuranceResponse, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("validating insurance")
	select {
	case <-time.After(time.Second * time.Duration(cmd.DelaySeconds)):
		logger.Info("validation received. unblocking")
		return &messages.ValidateInsuranceResponse{
			OrderID:    cmd.OrderID,
			CustomerID: cmd.CustomerID,
		}, nil
	}
}
