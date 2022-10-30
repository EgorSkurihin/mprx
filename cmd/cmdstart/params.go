package cmdstart

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

type Start struct {
	ProxyPort   int
	WebSitePort int

	ConfigUpdateInterval time.Duration
	ConfigPath           string

	LogPath               string
	MetricsPath           string
	MetricsUpdateInterval time.Duration
}

func (params *Start) SetupFlags(cmd *cobra.Command) {
	flags := cmd.Flags()

	flags.IntVar(&params.ProxyPort, "proxy-port", 3305, "http server port")
	flags.IntVar(&params.WebSitePort, "website-port", 4321, "http server port")

	flags.StringVar(&params.ConfigPath, "config-path", "", "url to feeds config file (required)")

	flags.DurationVar(&params.ConfigUpdateInterval, "config-update-interval", time.Minute, "config update interval")

	if err := cmd.MarkFlagRequired("config-path"); err != nil {
		panic(err)
	}

	flags.StringVar(&params.LogPath, "log", "", "log filename; use '-' or omit flag for stderr")

	flags.StringVar(&params.MetricsPath, "metrics-path", "", "path for file with metrics")
	flags.DurationVar(&params.MetricsUpdateInterval, "metrics-update-interval", 10*time.Second, "metrics update interval")
}

func (params *Start) Validate() error {
	if params.ConfigPath == "" {
		return fmt.Errorf(`the required parameter "%s" is empty`, "config-path")
	}
	return nil
}
