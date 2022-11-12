package main

import (
	"log"
	"os"

	"github.com/EgorSkurihin/mprx/cmd/cmdproxy"
	"github.com/EgorSkurihin/mprx/cmd/cmdstart"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := cobra.Command{
		SilenceErrors: true,
		SilenceUsage:  true,
		Use:           "mprx",
	}
	rootCmd.AddCommand(
		cmdstart.NewCmd(),
		cmdproxy.NewCmd(),
	)

	exitCode := 0
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err.Error())
		exitCode = 1
	}
	os.Exit(exitCode)
}
