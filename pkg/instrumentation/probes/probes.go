package probes

import (
	"context"
	"github.com/temporalio/temporal-demo-infra/pkg/instrumentation/log"
	"net/http"
	"time"
)

type Probes struct {
	srv *http.Server
}

func (p *Probes) Start(ctx context.Context) error {
	log.GetLogger(ctx).Info("starting probes server on port 8080")
	return p.srv.ListenAndServe()
}

func (p *Probes) Shutdown(ctx context.Context) {
	if err := p.srv.Shutdown(ctx); err != nil {
		log.GetLogger(ctx).Error("error shutting down probes", log.Fields{"err": err})
	}
}

// NewProbes creates a Probes which should eventually be
// legitimate and include a /startup probe. Some day :(
func NewProbes(ctx context.Context) (*Probes, error) {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/readiness", readinessHandler)

	srv := &http.Server{
		Handler:      mux,
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return &Probes{
		srv: srv,
	}, nil
}

func readinessHandler(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
}

func healthHandler(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
}
