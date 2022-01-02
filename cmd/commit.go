package cmd

import (
	"github.com/pa/dotman/utils"
	"github.com/spf13/cobra"
)

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:                "commit",
	Short:              "Record changes to the repository",
	Long:               `Record changes to the repository`,
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		options := append([]string{utils.GitDir, utils.WorkTree, "commit"}, args...)
		utils.GitCommand(true, "", options...)
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// commitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// commitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
