package main

import (
	"github.com/spf13/cobra"
	"github.com/unmango/go/cli"
	protocmd "github.com/unstoppablemango/go-protocmd/pkg"
)

var rootCmd = &cobra.Command{
	Use:   "protocmd",
	Short: "Convert specifications to commandline arguments",
	Run: func(cmd *cobra.Command, args []string) {
		if err := protocmd.ListenAndServe(); err != nil {
			cli.Fail(err)
		}
	},
}

func Execute() error {
	return rootCmd.Execute()
}
