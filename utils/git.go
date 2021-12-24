package utils

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

type NotInstalled struct {
	message string
	error
}

func (e *NotInstalled) Error() string {
	return e.message
}

func GitCommand(args ...string) (*exec.Cmd, error) {
	gitExe, err := exec.LookPath("git")
	if err != nil {
		if errors.Is(err, exec.ErrNotFound) {
			programName := "git"
			return nil, &NotInstalled{
				message: fmt.Sprintf("unable to find git executable in PATH; please install %s before retrying", programName),
				error:   err,
			}
		}
		return nil, err
	}
	return exec.Command(gitExe, args...), nil
}

func GitCommandRun(args ...string) {
	gitCmd, gitCmdErr := GitCommand(args...)
	if gitCmdErr != nil {
		fmt.Print(gitCmdErr)
	}
	// for interactive command line
	gitCmd.Stdout = os.Stdout
	gitCmd.Stdin = os.Stdin
	gitCmd.Stderr = os.Stderr
	_ = gitCmd.Run()
}
