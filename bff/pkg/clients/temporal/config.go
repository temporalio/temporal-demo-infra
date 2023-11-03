package temporal

type Config struct {
	HostPort     string
	Namespace    string
	CertFilePath string `split_words:"true"`
	KeyFilePath  string `split_words:"true"`
	CloudCertPem string
	CloudCertKey string
}

func (c *Config) Prefix() string {
	return "temporal"
}
