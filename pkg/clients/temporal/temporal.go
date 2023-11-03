package temporal

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/temporalio/temporal-demo-infra/pkg/instrumentation/log"
	"github.com/uber-go/tally/v4"
	sdkclient "go.temporal.io/sdk/client"
	sdktally "go.temporal.io/sdk/contrib/tally"
	"logur.dev/logur"
	"net"
	"os"
)

func GetIdentity(taskQueue string) string {
	return fmt.Sprintf("%d@%s@%s", os.Getpid(), getHostName(), taskQueue)

}

func NewMetricsHandler(scope tally.Scope) sdkclient.MetricsHandler {
	return sdktally.NewMetricsHandler(scope)
}

type Clients struct {
	Client          sdkclient.Client
	NamespaceClient sdkclient.NamespaceClient
	Config          *Config
	ClientOptions   sdkclient.Options
	logger          logur.Logger
}

func (c *Clients) Close() error {
	if c.Client != nil {
		c.Client.Close()
	}
	if c.NamespaceClient != nil {
		c.NamespaceClient.Close()
	}
	return nil
}

func getHostName() string {
	hostName, err := os.Hostname()
	if err != nil {
		hostName = "Unknown"
	}
	return hostName
}

// NewClients creates the temporal and temporalNamespace clients
func NewClients(ctx context.Context, opts ...Option) (*Clients, error) {
	result := &Clients{
		ClientOptions: sdkclient.Options{},
	}
	opts = append([]Option{
		WithLogger(log.GetLogger(ctx)),
	}, opts...)
	for _, o := range opts {
		o(result)
	}
	logFields := log.Fields{
		"cfg_namespace":                     result.Config.Namespace,
		"cfg_hostport":                      result.Config.HostPort,
		"cfg_tls_disable_host_verification": result.Config.TlsDisableHostVerification,
		"cfg_cloud_cert_pem":                result.Config.CloudCertPem != "",
		"cfg_cloud_cert_key":                result.Config.CloudCertKey != "",
		"cfg_cert_file_path":                result.Config.CertFilePath,
		"cfg_key_file_path":                 result.Config.KeyFilePath,
	}
	logger := log.WithFields(log.GetLogger(ctx), log.Fields{})
	logger.Info("temporal_config")
	if result.logger != nil {
		logger = log.WithFields(result.logger, logFields)
	}
	// map
	if result.ClientOptions.HostPort == "" {
		result.ClientOptions.HostPort = result.Config.HostPort
	}
	if result.ClientOptions.Namespace == "" {
		result.ClientOptions.Namespace = result.Config.Namespace
	}
	if result.ClientOptions.Identity == "" {
		// same behavior as SDK
		result.ClientOptions.Identity = GetIdentity("")
	}
	// tls
	tlsConfig, err := getTLSConfig(ctx, logger, result.Config)
	if err != nil {
		logger.Error("failed to create tls config", log.Fields{"err": err})
		return nil, fmt.Errorf("failed to create tls config %w", err)
	}

	logger = log.WithFields(logger, log.Fields{
		"client_options_namespace": result.ClientOptions.Namespace,
		"client_options_address":   result.ClientOptions.HostPort,
	})
	result.ClientOptions.Logger = logur.LoggerToKV(logger)
	result.ClientOptions.ConnectionOptions = sdkclient.ConnectionOptions{TLS: tlsConfig}
	tc, err := sdkclient.Dial(result.ClientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to new temporal client %w", err)
	}
	logger.Info("created temporal client")
	result.Client = tc

	// TODO: NamespaceClient creation
	//if result.NamespaceClient, err = sdkclient.NewNamespaceClient(result.ClientOptions); err != nil {
	//	return nil, fmt.Errorf("failed to new namespace client %w", err)
	//}
	return result, nil
}

func getTLSConfig(_ context.Context, logger log.Logger, config *Config) (*tls.Config, error) {

	if config.CertFilePath == "" && config.CloudCertPem == "" {
		logger.Info("TLS not configured. skipping.")
		return nil, nil
	}
	serverName, _, parseErr := net.SplitHostPort(config.HostPort)
	if parseErr != nil {
		return nil, fmt.Errorf("failed to split hostport %s: %w", config.HostPort, parseErr)
	}
	logger = log.WithFields(logger, log.Fields{"server_name": serverName})
	var cert tls.Certificate
	if config.CertFilePath != "" && config.KeyFilePath != "" {
		var err error
		logger.Info("configuring tls")
		cert, err = tls.LoadX509KeyPair(config.CertFilePath, config.KeyFilePath)
		if err != nil {
			return nil, fmt.Errorf("failed to load TLS from files: %w", err)
		}
	} else if config.CloudCertPem != "" && config.CloudCertKey != "" {
		var err error
		logger.Info("configuring tls using pem data")
		cert, err = tls.X509KeyPair([]byte(config.CloudCertPem), []byte(config.CloudCertKey))
		if err != nil {
			return nil, fmt.Errorf("failed to load Cloud TLS from data: %w", err)
		}
	}
	return &tls.Config{
		Certificates:       []tls.Certificate{cert},
		ServerName:         serverName,
		InsecureSkipVerify: config.TlsDisableHostVerification,
	}, nil
}
