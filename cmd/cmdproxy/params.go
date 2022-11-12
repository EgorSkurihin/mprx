package cmdproxy

import (
	"time"

	"github.com/spf13/cobra"
)

type Proxy struct {
	ProxyAddr   string
	MongoDBAddr string

	ConfigUpdateInterval time.Duration
	ConfigPath           string

	LogPath               string
	MetricsPath           string
	MetricsUpdateInterval time.Duration
}

func (params *Proxy) SetupFlags(cmd *cobra.Command) {
	flags := cmd.Flags()

	flags.StringVar(&params.ProxyAddr, "proxy-addr", "localhost:27018", "proxy address")
	flags.StringVar(&params.MongoDBAddr, "mongo-addr", "localhost:27017", "mongo-db address")

	flags.StringVar(&params.ConfigPath, "config-path", "config.yaml", "url to feeds config file (required)")

	flags.DurationVar(&params.ConfigUpdateInterval, "config-update-interval", time.Minute, "config update interval")

	flags.StringVar(&params.LogPath, "log", "", "log filename; use '-' or omit flag for stderr")

	flags.StringVar(&params.MetricsPath, "metrics-path", "", "path for file with metrics")
	flags.DurationVar(&params.MetricsUpdateInterval, "metrics-update-interval", 10*time.Second, "metrics update interval")
}

func (params *Proxy) Validate() error {
	return nil
}
