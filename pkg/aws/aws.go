package aws

import (
	"os"
	"gopkg.in/AlecAivazis/survey.v1"
	"github.com/go-ini/ini"
	"fmt"
	"path/filepath"
	"github.com/gosuri/uitable"
)

const (
	envProfile   = "AWS_PROFILE"
	envAccessKey = "AWS_ACCESS_KEY_ID"
	envSecretKey = "AWS_SECRET_ACCESS_KEY"
)

func promptUserAWSProfile(current string, opts []string) (string, error) {
	var qs = []*survey.Question{
		{
			Name: "profile",
			Prompt: &survey.Select{
				Message: "Choose an aws profile:",
				Options: opts,
				Default: current,
			},
		},
	}

	answers := struct {
		AWSProfile string `survey:"profile"`
	}{}

	err := survey.Ask(qs, &answers)
	if err != nil {
		return "", err
	}
	return answers.AWSProfile, nil
}

type credentials struct {
	Profile           string
	AccessKeyID       string
	SecretAccessKeyID string
}

func (c credentials) envString(justProfile bool) string {
	if justProfile {
		return fmt.Sprintf("unset %v; unset %v; export %v=%v", envAccessKey, envSecretKey, envProfile, c.Profile)
	}
	return fmt.Sprintf("unset %v; export %v=%v; export %v=%v", envProfile, envAccessKey, c.AccessKeyID, envSecretKey, c.SecretAccessKeyID)
}

func readAWSCredentials() ([]credentials, error) {
	f, err := ini.Load(filepath.Join(os.Getenv("HOME"), ".aws/credentials"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	s := f.Sections()

	var creds = make([]credentials, 0)

	for _, v := range s {
		if v.Name() != "DEFAULT" {
			id, err := v.GetKey("aws_access_key_id")
			if err != nil {
				return creds, fmt.Errorf("error reading aws credentials file: %v", err)
			}

			secret, err := v.GetKey("aws_secret_access_key")
			if err != nil {
				return creds, fmt.Errorf("error reading aws credentials file: %v", err)
			}
			creds = append(creds, credentials{
				Profile:           v.Name(),
				AccessKeyID:       id.String(),
				SecretAccessKeyID: secret.String(),
			})
		}
	}
	return creds, nil
}

func getProfileNames(creds []credentials) []string {
	s := make([]string, 0)
	for _, v := range creds {
		s = append(s, v.Profile)
	}
	return s
}

func formatCredsList(creds []credentials) string {
	current := getCurrentProfile(creds)

	table := uitable.New()
	if current != "" {
		table.AddRow("CURRENT", "PROFILE", "ACCESS_KEY_ID")
	}
	table.AddRow("PROFILE", "ACCESS_KEY_ID")


	for _, v := range creds {
		if current != "" {
			if v.Profile == current {
				table.AddRow("*", v.Profile, v.AccessKeyID)
			} else {
				table.AddRow("", v.Profile, v.AccessKeyID)
			}
		} else {
			table.AddRow(v.Profile, v.AccessKeyID)
		}

	}
	return table.String()
}

func getCurrentProfile(creds []credentials) string {
	env := credentials{
		Profile:           os.Getenv(envProfile),
		AccessKeyID:       os.Getenv(envAccessKey),
		SecretAccessKeyID: os.Getenv(envSecretKey),
	}
	if env.Profile != "" {
		return env.Profile
	}

	if env.AccessKeyID != "" {
		for _, v := range creds {
			if v.AccessKeyID == env.AccessKeyID {
				return v.Profile
			}
		}
	}

	return ""
}
