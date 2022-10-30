package cmdstart

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	var params Start

	cmd := cobra.Command{
		Use:   "start",
		Short: "Start proxy server for mysql db and website watch logs and dashboards",
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

func runStartCmd(params *Start) error {
	facade, err := newStartFacade(params)
	if err != nil {
		return err
	}
	return facade.startServer()
}
