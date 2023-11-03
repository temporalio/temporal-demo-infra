package middleware

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/temporalio/temporal-demo-infra/bff/pkg/instrumentation/log"
	"logur.dev/logur"
)

// Logger is a middleware that logs the start and end of each request, along
// with some useful data about what was requested, what the response status was,
// and how long it took to return.
// shamelessly stolen from https://github.com/treastech/logger
func Logger(l logur.Logger) func(next http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = log.WithLogger(ctx, l)
			r = r.WithContext(ctx)
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			t1 := time.Now()
			defer func() {
				l.Info("requested",
					logur.Fields{
						"proto":      r.Proto,
						"path":       r.URL.Path,
						"duration":   time.Since(t1),
						"status":     ww.Status(),
						"size":       ww.BytesWritten(),
						"request_id": middleware.GetReqID(r.Context()),
					})
			}()

			next.ServeHTTP(ww, r)
		}
		return http.HandlerFunc(fn)
	}
}
