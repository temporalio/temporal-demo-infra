package temporal

import (
	sdkclient "go.temporal.io/sdk/client"
	"logur.dev/logur"
)

type Option func(c *Clients)

func WithConfig(cfg *Config) Option {
	return func(c *Clients) {
		c.Config = cfg
	}
}
func WithLogger(l logur.Logger) Option {
	return func(c *Clients) {
		c.logger = l
	}
}
func WithOptions(opts sdkclient.Options) Option {
	return func(c *Clients) {
		c.ClientOptions = opts
	}
}
