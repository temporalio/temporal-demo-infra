package clients

import (
	"github.com/hashicorp/go-multierror"
	"github.com/temporalio/temporal-demo-infra/bff/internal/clients/temporal"
	"logur.dev/logur"
)

type Option func(*Clients)

func WithTemporal(t *temporal.Clients, err error) Option {
	return func(c *Clients) {
		c.temporal = t
		c.clientErrors = multierror.Append(c.clientErrors, err)
	}
}
func WithLogger(l logur.Logger) Option {
	return func(c *Clients) {
		c.logger = l
	}
}
