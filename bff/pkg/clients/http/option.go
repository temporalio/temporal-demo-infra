package http

import (
	"context"
	"github.com/gojek/heimdall/v7"
	"github.com/gojek/heimdall/v7/httpclient"
	"logur.dev/logur"
	"net/http"
	"time"
)

type Option func(context.Context, *Client)

func WithBackoff(backoff heimdall.Backoff) Option {
	return func(ctx context.Context, c *Client) {
		c.opts = append(c.opts, httpclient.WithRetrier(heimdall.NewRetrier(backoff)))
	}
}
func WithTimeout(t time.Duration) Option {
	return func(ctx context.Context, c *Client) {
		c.opts = append(c.opts, httpclient.WithHTTPTimeout(t))
	}
}

func WithRetryCount(count int) Option {
	return func(ctx context.Context, c *Client) {
		c.opts = append(c.opts, httpclient.WithRetryCount(count))
	}
}
func WithLogger(logger logur.Logger) Option {
	return func(ctx context.Context, c *Client) {
		c.plugins = append(c.plugins, NewRequestLogger(logger))
	}
}
func WithPlugin(plugin heimdall.Plugin) Option {
	return func(ctx context.Context, c *Client) {
		c.plugins = append(c.plugins, plugin)
	}
}
func WithInnerHTTPClient(cli *http.Client) Option {
	return func(ctx context.Context, c *Client) {
		c.httpClient = cli
	}
}
func WithConfig(cfg *Config) Option {
	return func(ctx context.Context, c *Client) {
		c.cfg = cfg
	}
}
