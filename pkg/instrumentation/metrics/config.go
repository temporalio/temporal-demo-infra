package metrics

type Config struct {
	ListenAddress string
	TimerType     string
	ScopePrefix   string
}

func (c *Config) Prefix() string {
	return "metrics"
}
