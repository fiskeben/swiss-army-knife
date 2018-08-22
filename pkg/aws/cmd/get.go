package cmd

import (
	"io"

	"github.com/spf13/cobra"
)

func newAwsGetCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get [FLAGS] credentials",
		Short: "get current aws credentials",
	}

	cmd.AddCommand(newAwsGetCredentialsCmd(out))
	return cmd
}
