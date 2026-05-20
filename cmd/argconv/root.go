package main

import (
	"github.com/spf13/cobra"
	"github.com/unmango/go/cli"
	cmdv1alpha1 "github.com/unstoppablemango/go-protocmd/gen/dev/unmango/cmd/v1alpha1"
	"github.com/unstoppablemango/go-protocmd/pkg/conv"
	"github.com/unstoppablemango/go-protocmd/pkg/conv/codec"
	"github.com/unstoppablemango/go-protocmd/pkg/log"
)

var rootCmd = &cobra.Command{
	Use:   "argconv",
	Short: "Convert specifications to commandline arguments",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		log.WriteTo(cmd.OutOrStderr())
	},
	Run: func(cmd *cobra.Command, args []string) {
		var req cmdv1alpha1.ArgsRequest
		d := codec.ArgsRequest.NewDecoder(cmd.InOrStdin())
		if err := d.Decode(&req); err != nil {
			cli.Fail(err)
		}

		a, err := conv.Args(req.GetSpec().ProtoReflect())
		if err != nil {
			cli.Fail(err)
		}

		res := &cmdv1alpha1.ArgsResponse_builder{Args: a}
		enc := codec.ArgsResponse.NewEncoder(cmd.OutOrStdout())
		if err := enc.Encode(res.Build()); err != nil {
			cli.Fail(err)
		}
	},
}

func Execute() error {
	return rootCmd.Execute()
}
