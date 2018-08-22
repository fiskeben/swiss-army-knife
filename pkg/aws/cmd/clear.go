package cmd

import (
	"io"
	"github.com/spf13/cobra"
)

func newAwsClearCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "clear [FLAGS] credentials",
		Short: "get a string copied into clipboard to clear all environment credentials",
	}

	cmd.AddCommand(newAwsClearCredentialsCmd(out))
	return cmd
}