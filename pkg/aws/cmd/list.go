package cmd

import (
	"io"
	"github.com/spf13/cobra"
)

func newAwsListCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list [FLAGS] credentials",
		Aliases: []string{"ls"},
		Short:   "list aws credentials",
	}

	cmd.AddCommand(newAwsListCredentialsCmd(out))
	return cmd
}
