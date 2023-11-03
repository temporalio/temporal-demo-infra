package provider_aws

//
//import (
//	"context"
//	"fmt"
//	"github.com/temporalio/temporal-demo-infra/pkg/messages"
//	"go.temporal.io/sdk/activity"
//	sdkclient "go.temporal.io/sdk/client"
//	"time"
//)
//
//var TypeHandlers *Handlers
//
//type Handlers struct {
//	temporalClient sdkclient.Client
//}
//
//func NewHandlers(temporalClient sdkclient.Client) *Handlers {
//	return &Handlers{temporalClient: temporalClient}
//}
//
//// ProvisionFoundationResources accepts a command and signals the Order of fulfillment after N seconds
//func (h *Handlers) ProvisionFoundationResources(ctx context.Context, cmd *messages.ProvisionFoundationResourcesRequest) (*messages.ProvisionFoundationResourcesResponse, error) {
//	logger := activity.GetLogger(ctx)
//	logger.Info(
//		"provisioning foundation resources",
//		"applicationId",
//		cmd.Region,
//		"prescriptionId",
//		cmd.Profile,
//	)
//	event := &messages.ProvisionFoundationResourcesResponse{
//		ApplicationID:      cmd.Profile,
//		CompletionDateTime: time.Now().Format("2006/01/02"),
//	}
//	if cmd.DelaySeconds == 0 {
//		// synchronous reply
//		return event, nil
//	}
//
//	// schedule to send event non-blocking
//	go (func(cmd *messages.ProvisionFoundationResourcesRequest) {
//		select {
//		case <-time.After(time.Duration(cmd.DelaySeconds) * time.Second):
//			logger.Info("signaling prescription fulfillment")
//			if err := h.temporalClient.SignalWorkflow(context.Background(), cmd.Region, "", messages.MessageName(event), event); err != nil {
//				fmt.Println(fmt.Errorf("failed to trigger event: %w", err))
//			}
//		}
//	})(cmd)
//	return nil, nil
//}
//func (h *Handlers) DestroyResources(ctx context.Context, cmd *messages.ProvisionFoundationResourcesResponse) error {
//	logger := activity.GetLogger(ctx)
//	logger.Info("destroying resources", "applicationID", cmd.ApplicationID)
//	return nil
//}
