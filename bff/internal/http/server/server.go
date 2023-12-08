package server

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/hashicorp/go-multierror"
	temporalClient "github.com/temporalio/temporal-demo-infra/bff/internal/clients/temporal"
	"github.com/temporalio/temporal-demo-infra/bff/internal/http/provision"

	"github.com/go-chi/cors"
	"github.com/temporalio/temporal-demo-infra/bff/internal/http/app"
	"github.com/temporalio/temporal-demo-infra/bff/internal/http/middleware"
	"github.com/temporalio/temporal-demo-infra/bff/internal/http/routes"
	"github.com/temporalio/temporal-demo-infra/bff/pkg/instrumentation/log"
	"logur.dev/logur"
	"net/http"
	"net/http/httputil"
)

type Closeable interface {
	Close() error
}
type Server struct {
	temporal   *temporalClient.Clients
	logger     logur.Logger
	router     *chi.Mux
	cfg        *Config
	inner      *http.Server
	errors     *multierror.Error
	closeables []Closeable
	taskQueue  string
}

// NewServer creates a new server with options
// A new root router will be created if one is not provided
func NewServer(ctx context.Context, opts ...Option) (*Server, error) {
	defaultOpts := []Option{
		WithRouter(chi.NewRouter()),
		WithConfig(&Config{}),
	}
	opts = append(defaultOpts, opts...)

	s := &Server{
		errors:     &multierror.Error{},
		closeables: make([]Closeable, 0),
	}

	for _, o := range opts {
		o(s)
	}

	// all routes use this middleware
	s.router.Use(middleware.Logger(s.logger))
	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	s.router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"http://*", "https://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-TeamID", "Connection"},
		ExposedHeaders:   []string{"Link", "Connection", "Cache-Control", "Content-Type"},
		AllowCredentials: true,
		Debug:            true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	logger := log.GetLogger(ctx)
	logger.Info("registering routers")

	s.router.Mount(routes.GETApi.Raw, s.buildApiRouter(s.router))
	if s.errors.ErrorOrNil() != nil {
		return nil, s.errors.ErrorOrNil()
	}

	s.router.Get("/ping", pingHandler)
	s.router.Get("/health", healthHandler)
	s.router.Get("/readiness", readinessHandler)

	s.inner = &http.Server{
		Addr:    fmt.Sprintf(":%s", s.cfg.Port),
		Handler: s.router,
	}
	for _, r := range s.router.Routes() {
		log.GetLogger(ctx).Info("registered the route", log.Fields{"route": r.Pattern})
	}
	return s, nil
}
func (s *Server) appendError(err error) error {
	s.errors = multierror.Append(s.errors, err)
	return s.errors
}

func (s *Server) buildPublicRouter(r chi.Router) {
	var err error
	if s.cfg.IsServingUI {
		var appHandlers *app.Handlers
		if appHandlers, err = app.NewHandlers(
			app.WithGeneratedAppDirectory(s.cfg.GeneratedAppDir),
			app.WithTemporalClients(s.temporal),
			app.WithMountPath(routes.GETApp.Raw),
		); err != nil {
			_ = s.appendError(err)
			return
		}

		appHandlers.Register(r)
	}

}
func (s *Server) buildApiRouter(r chi.Router) chi.Router {

	provisionHandlers, err := provision.NewHandlers(provision.WithTemporalClients(s.temporal))
	if err != nil {
		_ = s.appendError(err)
		return nil
	}
	r.Post("/provision", provisionHandlers.POST)
	r.Patch("/provision", provisionHandlers.PATCH)
	r.Delete("/provision", provisionHandlers.DELETE)
	return r
}

// Start starts the server
func (s *Server) Start(ctx context.Context) error {
	log.GetLogger(ctx).Info("starting http server", log.Fields{
		"port": s.cfg.Port,
	})

	return s.inner.ListenAndServe()
	//return s.inner.ListenAndServeTLS("/Users/mnichols/certs/localhost/server.crt", "/Users/mnichols/certs/localhost/server.key")
}
func (s *Server) Shutdown(ctx context.Context) {
	if err := s.inner.Shutdown(ctx); err != nil {
		s.logger.Error("failed to shutdown gracefully", logur.Fields{"err": err})
	}
}
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
func pingHandler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		if i, werr := w.Write([]byte(fmt.Sprintf("pong but error %s", err.Error()))); werr != nil {
			fmt.Println("wrote ", i, " bytes", werr)
		}
		return
	}
	if i, werr := w.Write(dump); werr != nil {
		fmt.Println("wrote ", i, "bytes", werr)
	}
}
