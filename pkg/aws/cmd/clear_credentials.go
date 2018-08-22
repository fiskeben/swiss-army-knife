package cmd

import (
	"io"
	"github.com/spf13/cobra"
	"github.com/majestic-fox/swiss-army-knife/pkg/aws"
	"github.com/atotto/clipboard"
)

type awsClearCredentialsCmd struct {
	out io.Writer
}

func newAwsClearCredentialsCmd(out io.Writer) *cobra.Command {
	a := &awsClearCredentialsCmd{out: out}
	cmd := &cobra.Command{
		Use:     "credentials",
		Aliases: []string{"creds"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := a.run(); err != nil {
				return err
			}
			return nil
		},
	}
	return cmd
}

func (a *awsClearCredentialsCmd) run() error {
	clipboard.WriteAll(aws.ClearCredentials())
	return nil
}