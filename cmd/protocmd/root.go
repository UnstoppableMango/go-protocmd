package main

import (
	"github.com/spf13/cobra"
	"github.com/unmango/go/cli"
	protocmd "github.com/unstoppablemango/go-protocmd/pkg"
	"github.com/unstoppablemango/go-protocmd/pkg/log"
)

var rootCmd = &cobra.Command{
	Use:   "protocmd",
	Short: "Convert specifications to commandline arguments",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		log.WriteTo(cmd.OutOrStderr())
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := protocmd.ListenAndServe(); err != nil {
			cli.Fail(err)
		}
	},
}

func Execute() error {
	return rootCmd.Execute()
}
