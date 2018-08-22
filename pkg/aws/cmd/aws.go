package cmd

import (
	"io"
	"github.com/spf13/cobra"
	aws_env "github.com/majestic-fox/swiss-army-knife/pkg/aws/environment"
)

var (
	settings aws_env.EnvSettings
)

func NewAwsCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "aws [FLAGS] get|set",
		Short: "aws commands",
	}

	cmd.AddCommand(newAwsGetCmd(out))
	cmd.AddCommand(newAwsSetCmd(out))
	cmd.AddCommand(newAwsListCmd(out))
	cmd.AddCommand(newAwsClearCmd(out))

	flags := cmd.PersistentFlags()
	settings.AddFlags(flags)
	settings.Init(flags)

	return cmd
}