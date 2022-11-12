package cmdproxy

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	var params Proxy

	cmd := cobra.Command{
		Use:   "proxy",
		Short: "Start proxy server for mongodb db",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := params.Validate(); err != nil {
				return fmt.Errorf("fail to validate start flags: %w", err)
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return runStartCmd(&params)
		},
	}

	params.SetupFlags(&cmd)

	return &cmd
}

func runStartCmd(params *Proxy) error {
	facade, err := newProxyFacade(params)
	if err != nil {
		return err
	}
	return facade.start()
}
