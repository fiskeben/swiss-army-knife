package cmd

import (
	"io"
	"github.com/spf13/cobra"
	"github.com/majestic-fox/swiss-army-knife/pkg/aws"
	"fmt"
)

type awsGetCredentialsCmd struct {
	out io.Writer
}

func newAwsGetCredentialsCmd(out io.Writer) *cobra.Command {
	a := &awsGetCredentialsCmd{out: out}
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

func (a *awsGetCredentialsCmd) run() error {
	c, err := aws.GetCurrentCredentials()
	if err != nil {
		if settings.Quiet {
			fmt.Print("☠️")
			return nil
		}
	}
	if c == "" {
		fmt.Print("🤲")
	} else {
		fmt.Print(c)
	}
	return nil
}
