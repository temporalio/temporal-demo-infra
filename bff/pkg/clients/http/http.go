package http

import (
	"context"
	"fmt"
	"github.com/gojek/heimdall/v7"
	"github.com/gojek/heimdall/v7/httpclient"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Doer = heimdall.Doer

type Client struct {
	httpClient *http.Client
	inner      *httpclient.Client
	plugins    []heimdall.Plugin
	opts       []httpclient.Option
	cfg        *Config
}

func (c *Client) GetUnderlyingHTTPClient() *http.Client {
	return c.httpClient
}
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	return c.inner.Do(req)
}

// NewClient creates a new http client
func NewClient(ctx context.Context, opts ...Option) (*Client, error) {
	// TODO set TLS and Dialer timeouts
	//transport := &nethttp.Transport{}
	defaults := []Option{
		WithInnerHTTPClient(&http.Client{
			Transport: &http.Transport{},
		}),
		WithLogger(nil),
		WithBackoff(heimdall.NewExponentialBackoff(
			2*time.Millisecond,
			9*time.Millisecond,
			2,
			2*time.Millisecond,
		)),
		WithRetryCount(4),
		WithTimeout(1000 * time.Millisecond),
	}

	opts = append(defaults, opts...)
	httpClient := &http.Client{
		//Transport: transport,
	}
	result := &Client{
		httpClient: httpClient,
		inner:      nil,
		opts:       []httpclient.Option{httpclient.WithHTTPClient(httpClient)},
		plugins:    make([]heimdall.Plugin, 0),
	}
	for _, o := range opts {
		o(ctx, result)
	}

	result.inner = httpclient.NewClient(result.opts...)
	for _, p := range result.plugins {
		result.inner.AddPlugin(p)
	}
	return result, nil
}

func MustNewClient(ctx context.Context, opts ...Option) *Client {
	c, err := NewClient(ctx, opts...)
	if err != nil {
		panic(fmt.Errorf("failed to new http client %w", err).Error())
	}
	return c
}
func NewRequest(ctx context.Context, method string, url string, body io.Reader, params ...RequestParams) (*http.Request, error) {
	r, err := http.NewRequestWithContext(ctx, method, url, body)
	for _, p := range params {
		for k, v := range p.Header {
			r.Header[k] = v
		}
		for k, v := range p.Query {
			q := r.URL.Query()
			q[k] = v
			r.URL.RawQuery = q.Encode()
		}
	}
	return r, err
}

// RequestParams are customizable, request-scoped params for passing into explicit http clients
type RequestParams struct {
	Header http.Header
	Query  url.Values
}
