package http

import (
	"context"
	"github.com/temporalio/temporal-demo-infra/bff/pkg/instrumentation/log"
	"net/http/httputil"
	"time"

	"net/http"
)

type ctxKey string

const reqTimeCtxKey ctxKey = "reqTime"

type requestLogger struct {
	logger log.Logger
}

func getReqTime(ctx context.Context) (time.Time, bool) {
	tr, ok := ctx.Value(reqTimeCtxKey).(time.Time)
	if !ok {
		return time.Now(), ok
	}
	return tr, ok
}

// NewRequestLogger is Heimdall plugin for instrumentation.Logger
func NewRequestLogger(logger log.Logger) heimdall.Plugin {
	if logger != nil {
		logger = log.WithFields(logger, log.Fields{"client": "http_client"})
	}
	return &requestLogger{
		logger: logger,
	}
}

func (rl *requestLogger) OnRequestStart(req *http.Request) {
	ctx := context.WithValue(req.Context(), reqTimeCtxKey, time.Now())
	logger := rl.logger
	if rl.logger == nil {
		logger = log.WithFields(log.GetLogger(ctx), log.Fields{"client": "http_client"})
	}
	logger = log.WithFields(logger, log.Fields{
		"method": req.Method,
		"url":    req.URL.String(),
	})
	logger.Debug("request_start")
	ctx = log.WithLogger(ctx, logger)
	*req = *(req.WithContext(ctx))
}

func (rl *requestLogger) OnRequestEnd(req *http.Request, res *http.Response) {
	ctx := req.Context()
	logger := log.GetLogger(ctx)
	reqDuration := getRequestDuration(ctx) / time.Millisecond
	statusCode := res.StatusCode
	if statusCode > 399 {
		var err error
		var resDump []byte
		var reqDump []byte
		resDump, err = httputil.DumpResponse(res, true)
		if err != nil {
			logger.Error("failed to dump response", log.Fields{log.TagError: err})
		}
		reqDump, err = httputil.DumpRequestOut(req, true)
		if err != nil {
			logger.Error("failed to dump request", log.Fields{log.TagError: err})
		}
		logger.Debug("error response received", log.Fields{
			"request_dump":  reqDump,
			"response_dump": resDump,
		})
	}
	logger.Debug("request_end", log.Fields{
		"duration_ms": reqDuration,
		"status_code": statusCode,
	})
}

func (rl *requestLogger) OnError(req *http.Request, err error) {
	ctx := req.Context()
	logger := log.GetLogger(ctx)
	reqDuration := getRequestDuration(ctx) / time.Millisecond
	logger.Error("request_end", log.Fields{
		log.TagError:  err,
		"duration_ms": reqDuration,
	})
}
func getRequestDuration(ctx context.Context) time.Duration {
	now := time.Now()
	startTime, ok := getReqTime(ctx)
	if !ok {
		return 0
	}

	return now.Sub(startTime)
}
