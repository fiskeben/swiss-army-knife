package cmd

import (
	"io"
	"github.com/spf13/cobra"
	"github.com/majestic-fox/swiss-army-knife/pkg/aws"
	"fmt"
)

type awsListCredentialsCmd struct {
	out io.Writer
}

func newAwsListCredentialsCmd(out io.Writer) *cobra.Command {
	a := &awsListCredentialsCmd{out: out}
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

func (a *awsListCredentialsCmd) run() error {
	list, err := aws.ListCredentials()
	if err != nil {
		return err
	}
	fmt.Print(list)
	return nil
}
