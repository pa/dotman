package cmd

import (
	"github.com/pa/dotman/utils"
	"github.com/spf13/cobra"
)

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:                "pull",
	Short:              "Fetch from and merge with another repository or a local branch",
	Long:               `Fetch from and merge with another repository or a local branch`,
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		options := append([]string{utils.GitDir, utils.WorkTree, "pull"}, args...)
		utils.GitCommand(true, "", options...)
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pullCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pullCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
