package utils

import (
	"errors"
	"fmt"
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
			// if runtime.GOOS == "windows" {
			//  programName = "Git for Windows"
			// }
			return nil, &NotInstalled{
				message: fmt.Sprintf("unable to find git executable in PATH; please install %s before retrying", programName),
				error:   err,
			}
		}
		return nil, err
	}
	return exec.Command(gitExe, args...), nil
}
