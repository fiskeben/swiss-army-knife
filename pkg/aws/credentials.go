package aws

import "fmt"

func ListCredentials() (string, error) {
	c, err := readAWSCredentials()
	if err != nil {
		return "", fmt.Errorf("failed to read credentials file: %v", err)
	}
	return formatCredsList(c), nil
}

func SetCredentials(justProfile bool) (string, error) {
	c, err := readAWSCredentials()
	if err != nil {
		return "", fmt.Errorf("failed to read credentials file: %v", err)
	}

	newProfile, err := promptUserAWSProfile(getCurrentProfile(c), getProfileNames(c))
	if err != nil {
		return "", fmt.Errorf("something went wrong when prompting the user: %v", err)
	}

	for _, v := range c {
		if v.Profile == newProfile {
			return v.envString(justProfile), nil
		}
	}

	return "", fmt.Errorf("🤭 something went wrong while setting the credentials")
}

func GetCurrentCredentials() (string, error) {
	c, err := readAWSCredentials()
	if err != nil {
		return "", fmt.Errorf("failed to read credentials file: %v", err)
	}

	return getCurrentProfile(c), nil
}

func ClearCredentials() string {
	return fmt.Sprintf("unset %v; unset %v; unset %v", envProfile, envAccessKey, envSecretKey)
}