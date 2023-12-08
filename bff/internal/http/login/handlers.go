package login

import (
	"context"
	"fmt"
	"github.com/temporalio/temporal-demo-infra/bff/internal/clients/temporal"
	"github.com/temporalio/temporal-demo-infra/bff/internal/http/encoding"
	"github.com/temporalio/temporal-demo-infra/bff/pkg/instrumentation/log"
	"net/http"
)

type SessionStarter interface {
	StartSession(ctx context.Context, email string) error
	GenerateToken(ctx context.Context, email string) (string, error)
}

type Handlers struct {
	temporal      *temporal.Clients
	encryptionKey string
	password      string
	session       SessionStarter
}

func NewHandlers(opts ...Option) (*Handlers, error) {
	h := &Handlers{}
	for _, opt := range opts {
		opt(h)
	}
	if h.temporal == nil {
		return nil, fmt.Errorf("temporal client required and missing")
	}
	if h.session == nil {
		return nil, fmt.Errorf("authenticator is required")
	}
	return h, nil
}

func (h *Handlers) POST(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := log.GetLogger(ctx)
	u := &LoginPOST{}
	if err := encoding.DecodeJSONBody(w, r, u); err != nil {
		logger.Error("failed to authenticate", log.Fields{log.TagError: err})
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// primitive singular password for the whole world
	if u.Password != h.password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if err := h.session.StartSession(ctx, u.Email); err != nil {
		logger.Error("failed to start session", log.Fields{log.TagError: err, "email": u.Email})
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	token, err := h.session.GenerateToken(ctx, u.Email)
	if err != nil {
		logger.Error("failed to generate token", log.Fields{log.TagError: err, "email": u.Email})
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	result := &LoginResponse{
		Email: u.Email,
		Token: token,
	}
	if err := encoding.EncodeJSONResponseBody(w, result, http.StatusOK); err != nil {
		logger.Error("failed to write response", log.Fields{log.TagError: err, "email": u.Email})
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
