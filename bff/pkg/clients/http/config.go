package http

type Config struct {
}

func (c *Config) Prefix() string {
	return "httpclient"
}
