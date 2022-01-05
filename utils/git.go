package utils

import (
	"os"
	"strings"
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

// Check if repo already up to date
func IsRepoUpToDate(repoPath string) bool {
	isUpToDate, _ := GitCommand(false,
		repoPath,
		"diff",
		"origin/HEAD",
		"HEAD",
		"--name-only",
	)

	if len(isUpToDate) > 0 {
		return false
	}
	return true
}

func GetDeletedOrModifiedFiles(repoPath string, currentCommitID []byte, isDeletedFiles bool) []string {
	var diffFilter string
	if isDeletedFiles {
		diffFilter = "D"
	} else {
		diffFilter = "AM"
	}

	files, _ := GitCommand(false,
		repoPath,
		"show",
		RemoveRunes(string(currentCommitID))+"..HEAD",
		"--diff-filter="+diffFilter,
		"--name-only",
		"--no-commit-id",
	)
	return strings.Fields(string(files))
}
