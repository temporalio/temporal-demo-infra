package app

import (
	"fmt"
	"github.com/temporalio/temporal-demo-infra/bff/internal/clients/temporal"
	"github.com/temporalio/temporal-demo-infra/bff/internal/http/routes"
	"github.com/temporalio/temporal-demo-infra/bff/pkg/instrumentation/log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Router interface {
	Get(string, http.HandlerFunc)
}
type Handlers struct {
	temporal        *temporal.Clients
	inner           http.Handler
	generatedAppDir string
	mountPath       string
	fileSystem      http.FileSystem
}

func NewHandlers(opts ...Option) (*Handlers, error) {
	h := &Handlers{}
	for _, opt := range opts {
		opt(h)
	}
	if h.temporal == nil {
		return nil, fmt.Errorf("temporal client required and missing")
	}
	if h.fileSystem == nil {
		workDir, _ := os.Getwd()
		fmt.Println("using ", workDir)
		h.fileSystem = http.Dir(filepath.Join(workDir, h.generatedAppDir))
	}

	h.fileSystem = justFilesFilesystem{
		fs:               h.fileSystem,
		readDirBatchSize: 2,
	}

	pathPrefix := strings.TrimSuffix(h.mountPath, "/*")
	h.inner = http.StripPrefix(pathPrefix, http.FileServer(h.fileSystem))
	return h, nil
}

func (h *Handlers) GET(w http.ResponseWriter, r *http.Request) {
	fmt.Println("received ", r.URL.String())
	logger := log.GetLogger(r.Context())
	logger.Debug("rendering")
	h.inner.ServeHTTP(w, r)

}

func (h *Handlers) Register(router Router) {
	router.Get(routes.GETAppNoSlash.Raw, routes.GETAppNoSlash.Redirect.Handler().ServeHTTP)
	router.Get(routes.GETApp.Raw, h.GET)
}
