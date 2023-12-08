package api

import (
	"fmt"
	"github.com/temporalio/temporal-demo-infra/bff/internal/clients/temporal"
	"net/http"
)

type Router interface {
	Get(string, http.HandlerFunc)
}
type Handlers struct {
	temporal      *temporal.Clients
	encryptionKey string
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

func (h *Handlers) GET(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()
	//logger := log.GetLogger(r.Context())

	w.WriteHeader(http.StatusOK)
}
