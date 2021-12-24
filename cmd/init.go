package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pa/dotman/utils"
	"github.com/spf13/cobra"
	"github.com/tcnksm/go-gitconfig"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "clones dotfiles repo and installs plugins that you specified in cofig file",
	Long:  `clones dotfiles repo and installs plugins that you specified in cofig file.`,
	Run: func(cmd *cobra.Command, args []string) {
		Params := utils.PromptParams{
			Label:    "git repo url",
			ErrorMsg: "Please provide git repo url",
		}

		DotFilesRepoUrl := utils.PromptGetInput(Params)

		// Clone dotfiles bare repo
		utils.GitCommandRun(utils.GitDir,
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
		utils.GitCommandRun(utils.GitDir,
			utils.WorkTree,
			"config",
			"user.name",
			utils.GitUserName,
		)

		// set showUntrackedFiles to no
		utils.GitCommandRun(utils.GitDir,
			utils.WorkTree,
			"config",
			"--local",
			"status.showUntrackedFiles",
			"no",
		)

		// get branch name
		DefaultBranchName, _ := gitconfig.Entire("init.defaultbranch")

		// git list files in remote repo
		LsFiles := utils.GitCommandRun(utils.GitDir,
			"ls-tree",
			"-r",
			DefaultBranchName,
			"--name-only",
		)

		// convert to Array
		LsFilesArray := strings.Fields(LsFiles)

		// backup existing dotfiles
		if len(LsFilesArray) >= 0 {
			for _, filePath := range LsFilesArray {
				os.MkdirAll(utils.DotfileBackupDir+"/"+filepath.Dir(filePath), 0755)
				os.Rename(utils.HomeDir+"/"+filePath, utils.DotfileBackupDir+"/"+filePath)
			}
		}

		// git checkout dotfiles
		utils.GitCommandRun(utils.GitDir,
			utils.WorkTree,
			"checkout",
		)

		fmt.Print("Successfully Initialized")

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
