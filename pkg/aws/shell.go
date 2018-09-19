package aws

import (
	"fmt"
	"os"
	"os/user"
)

func openShell(env map[string]string) error {
	me, err := user.Current()
	if err != nil {
		return fmt.Errorf("failed to get current user: %v", err)
	}

	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current working directory: %v", err)
	}

	for k, v := range env {
		os.Setenv(k, v)
	}

	pa := os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		Dir:   cwd,
	}

	proc, err := os.StartProcess("/usr/bin/login", []string{"login", "-fpl", me.Username}, &pa)
	if err != nil {
		return fmt.Errorf("failed to start new shell: %v", err)
	}

	// Wait until user exits the shell
	state, err := proc.Wait()
	if err != nil {
		return fmt.Errorf("failed to wait for the user to exit the shell: %v", err)
	}

	fmt.Printf("<< Exited shell: %v\n", state.String())
	return nil
}
