package cli

import (
	"crypto/tls"
	"os"

	corev2 "github.com/echlebek/sensu-lite/api/core/v2"
	"github.com/echlebek/sensu-lite/cli/client"
	"github.com/echlebek/sensu-lite/cli/client/config"
	"github.com/echlebek/sensu-lite/cli/client/config/basic"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

// SensuCmdName name of the command
const SensuCmdName = "sensuctl"

// TypeError is the string returned in the event of an invalid type assertion
const TypeError = "TypeError"

// SensuCli is an instance of the Sensu command line client;
// encapsulates API client, logger & general configuration.
type SensuCli struct {
	Config config.Config
	Client client.APIClient
	Logger *logrus.Entry
	InFile *os.File
}

// New SensuCLI given persistent flags from command
func New(flags *pflag.FlagSet) *SensuCli {
	conf := basic.Load(flags)
	client := client.New(conf)
	logger := logrus.WithFields(logrus.Fields{
		"component": "cli-client",
	})

	tlsConfig := tls.Config{}

	if conf.TrustedCAFile() != "" {
		caCertPool, err := corev2.LoadCACerts(conf.TrustedCAFile())
		if err != nil {
			logger.Warn(err)
			logger.Warn("Trying to use the system's default CA certificates")
		}
		tlsConfig.RootCAs = caCertPool
	}

	tlsConfig.InsecureSkipVerify = conf.InsecureSkipTLSVerify()

	tlsConfig.BuildNameToCertificate()
	tlsConfig.CipherSuites = corev2.DefaultCipherSuites

	client.SetTLSClientConfig(&tlsConfig)

	return &SensuCli{
		Client: client,
		Config: conf,
		Logger: logger,
		InFile: os.Stdin,
	}
}
