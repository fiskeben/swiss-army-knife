package cmd

import (
	"io"

	"github.com/spf13/cobra"
)

func newAwsSetCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set [FLAGS] credentials",
		Short: "set aws credentials",
	}

	cmd.AddCommand(newAwsSetCredentialsCmd(out))
	return cmd
}
