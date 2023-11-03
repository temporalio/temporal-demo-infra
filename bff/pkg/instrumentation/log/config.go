package log

type Config struct {
	Environment string `default:"production"`
}

func (c *Config) Prefix() string {
	return "log"
}
