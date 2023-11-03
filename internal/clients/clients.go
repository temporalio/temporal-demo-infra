package clients

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-multierror"
	"github.com/temporalio/temporal-demo-infra/pkg/clients/temporal"

	"sync"

	"github.com/temporalio/temporal-demo-infra/pkg/instrumentation/log"

	"logur.dev/logur"
)

var once sync.Once
var oneClients *Clients

// Clients is a useful collection of clients for one-time initialization storage
// It should NOT be used as a collection to be passed around as a service locator.
type Clients struct {
	logger       logur.Logger
	clientErrors *multierror.Error
	temporal     *temporal.Clients
}

func (c *Clients) Temporal() *temporal.Clients {
	return c.temporal
}

func (c *Clients) Close() error {
	var errs *multierror.Error
	if c.temporal != nil {
		if err := c.temporal.Close(); err != nil {
			errs = multierror.Append(errs, err)
		}
	}

	return errs.ErrorOrNil()
}

// NewClients creates Clients dependencies
func NewClients(ctx context.Context, opts ...Option) (*Clients, error) {
	result := &Clients{
		clientErrors: &multierror.Error{},
	}
	for _, o := range opts {
		o(result)
	}

	if err := result.clientErrors.ErrorOrNil(); err != nil {
		return nil, fmt.Errorf("failed to new clients: %w", err)
	}
	return result, nil
}

// MustGetClients demands a clients instance with typical components
// configured by a top-level config
func MustGetClients(ctx context.Context, opts ...Option) *Clients {

	once.Do(func() {

		var err error
		logger := log.GetLogger(ctx)
		if oneClients, err = NewClients(ctx, opts...); err != nil {
			logger.Error("failed to get clients", logur.Fields{"err": err})

			panic(err)
		}

	})

	return oneClients
}
