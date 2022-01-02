package cmd

import (
	"github.com/pa/dotman/utils"
	"github.com/spf13/cobra"
)

// diffCmd represents the diff command
var diffCmd = &cobra.Command{
	Use:                "diff",
	Short:              "Show changes between commits, commit and working tree, etc",
	Long:               `Show changes between commits, commit and working tree, etc`,
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		options := append([]string{utils.GitDir, utils.WorkTree, "diff"}, args...)
		utils.GitCommand(true, "", options...)
	},
}

func init() {
	rootCmd.AddCommand(diffCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// diffCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// diffCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
