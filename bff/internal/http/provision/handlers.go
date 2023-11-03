package provision

import (
	"fmt"
	"github.com/temporalio/temporal-demo-infra/bff/internal/clients/temporal"
	"github.com/temporalio/temporal-demo-infra/bff/internal/http/encoding"
	"github.com/temporalio/temporal-demo-infra/bff/pkg/instrumentation/log"
	"github.com/temporalio/temporal-demo-infra/internal/orchestrations"
	"github.com/temporalio/temporal-demo-infra/pkg/messages"
	"github.com/teris-io/shortid"
	"go.temporal.io/sdk/client"
	"net/http"
)

type Handlers struct {
	temporal *temporal.Clients
}

func NewHandlers(opts ...Option) (*Handlers, error) {
	h := &Handlers{}
	for _, opt := range opts {
		opt(h)
	}
	if h.temporal == nil {
		return nil, fmt.Errorf("temporal client required and missing")
	}
	return h, nil
}

func (h *Handlers) POST(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	workflowID := shortid.MustGenerate()
	logger := log.WithFields(log.GetLogger(ctx), log.Fields{"workflow_id": workflowID})
	u := &ProvisionPOST{}
	if err := encoding.DecodeJSONBody(w, r, u); err != nil {
		logger.Error("failed to authenticate", log.Fields{log.TagError: err})
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	params := &messages.ProvisionApplicationRequest{
		ApplicationID:                 workflowID,
		RequesterID:                   shortid.MustGenerate(),
		TeamID:                        u.TeamID,
		ApplicationName:               u.ApplicationName,
		AuthorizerID:                  shortid.MustGenerate(),
		AuthorizationTimeoutSeconds:   120,
		DemoAuthorizationDelaySeconds: 0,
	}
	logger.Info("starting scenario")
	_, err := h.temporal.Client.ExecuteWorkflow(ctx, client.StartWorkflowOptions{
		ID:        workflowID,
		TaskQueue: "apps",
	}, orchestrations.TypeOrchestrations.ProvisionApplication, params)
	if err != nil {
		logger.Error("error starting", log.Fields{"err": err})
	}
	//var response *orchestrations.ProvisionApplicationState
	//if err := run.Get(ctx, &response); err != nil {
	//
	//	w.WriteHeader(http.StatusInternalServerError)
	//	return
	//}
	result := &ProvisionPOSTResponse{
		ApplicationName: u.ApplicationName,
		TeamID:          u.TeamID,
		AuthorizerID:    u.AuthorizerID,
		ApplicationID:   workflowID,
		WorkflowID:      workflowID,
	}
	if err := encoding.EncodeJSONResponseBody(w, result, http.StatusOK); err != nil {
		logger.Error("failed to write response", log.Fields{log.TagError: err, "appName": u.ApplicationName})
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func (h *Handlers) PATCH(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := log.GetLogger(ctx)
	u := &ProvisionPATCH{}
	if err := encoding.DecodeJSONBody(w, r, u); err != nil {
		logger.Error("failed to signal", log.Fields{log.TagError: err})
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	event := &messages.AuthorizationReceivedResponse{
		ApplicationID: u.ApplicationID,
		Region:        u.Region,
		Profile:       u.Profile,
		TeamID:        u.TeamID,
		IsApproved:    true,
	}
	logger.Info("approving scenario")
	err := h.temporal.Client.SignalWorkflow(ctx, u.WorkflowID, "", messages.MessageName(event), event)
	if err != nil {
		logger.Error("error signaling", log.Fields{"err": err})
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusAccepted)
}
