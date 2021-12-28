package utils

import (
	"os"
)

func GitCommand(isInteractive bool, currentDir string, args ...string) ([]byte, error) {

	var gitCmdOutput []byte
	var gitCmdError error

	gitCmd, gitCmdErr := RunCommand("git", args...)
	if gitCmdErr != nil {
		return nil, gitCmdErr
	}

	if currentDir != "" {
		gitCmd.Dir = currentDir
	}

	if isInteractive {
		// for interactive command line
		gitCmd.Stdout = os.Stdout
		gitCmd.Stdin = os.Stdin
		gitCmd.Stderr = os.Stderr
		gitCmdError = gitCmd.Run()
		if gitCmdError != nil {
			return nil, gitCmdError
		}
	} else {
		gitCmdOutput, gitCmdError = gitCmd.Output()
		if gitCmdError != nil {
			return gitCmdOutput, gitCmdError
		}
	}
	return gitCmdOutput, gitCmdError
}

func IsGitRepoDir(repoPath string) bool {
	_, err := GitCommand(false,
		repoPath,
		"rev-parse",
		"-r",
		"--is-inside-work-tree",
	)
	if err != nil {
		return false
	}
	return true
}
