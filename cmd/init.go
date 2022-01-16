package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pa/dotman/utils"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Clones dotfiles repo from remote git repository",
	Long:  `Clones dotfiles repo from remote git repository`,
	Run: func(cmd *cobra.Command, args []string) {
		Params := utils.PromptParams{
			Label:    "git repo url",
			ErrorMsg: "Please provide git repo url",
		}

		DotFilesRepoUrl := utils.PromptGetInput(Params)

		if !utils.IsGitRepoDir(utils.DotmanDir) {
			// Clone dotfiles bare repo
			utils.GitCommand(true,
				"",
				utils.GitDir,
				utils.WorkTree,
				"clone",
				"--bare",
				"--no-checkout",
				"--depth",
				"1",
				DotFilesRepoUrl,
				utils.DotmanDir,
			)

			// set git config username
			utils.GitCommand(true,
				"",
				utils.GitDir,
				utils.WorkTree,
				"config",
				"user.name",
				utils.GitUserName,
			)

			// set showUntrackedFiles to no
			utils.GitCommand(true,
				"",
				utils.GitDir,
				utils.WorkTree,
				"config",
				"--local",
				"status.showUntrackedFiles",
				"no",
			)

			// get branch name
			CurrentBranchName, _ := utils.GitCommand(false,
				"",
				utils.GitDir,
				utils.WorkTree,
				"branch",
				"--show-current",
			)

			// git list files in remote repo
			gitCmd, err := utils.GitCommand(false,
				"",
				utils.GitDir,
				"ls-tree",
				"-r",
				utils.RemoveRunes(string(CurrentBranchName)),
				"--name-only",
			)
			if err != nil {
				fmt.Print(err)
				os.Exit(1)
			}

			LsFiles := strings.Fields(string(gitCmd))

			// backup existing dotfiles
			if len(LsFiles) >= 0 {
				for _, filePath := range LsFiles {
					os.MkdirAll(utils.DotfileBackupDir+"/"+filepath.Dir(filePath), os.ModePerm)
					os.Rename(utils.HomeDir+"/"+filePath, utils.DotfileBackupDir+"/"+filePath)
				}
			}

			// git checkout dotfiles
			utils.GitCommand(true,
				"",
				utils.GitDir,
				utils.WorkTree,
				"checkout",
			)

			fmt.Print("Successfully Initialized")
		} else {
			fmt.Print("Already Initialized")
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
