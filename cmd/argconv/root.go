package main

import (
	"github.com/spf13/cobra"
	"github.com/unstoppablemango/go-protocmd/pkg/log"
)

var rootCmd = &cobra.Command{
	Use:   "argconv",
	Short: "Convert specifications to commandline arguments",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		log.WriteTo(cmd.OutOrStderr())
	},
	Run: func(cmd *cobra.Command, args []string) {
		// TODO
	},
}

func Execute() error {
	return rootCmd.Execute()
}
