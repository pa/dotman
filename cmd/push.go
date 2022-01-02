package cmd

import (
	"github.com/pa/dotman/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// pushCmd represents the push command
var pushCmd = &cobra.Command{
	Use:                "push",
	Short:              "Update remote refs along with associated objects",
	Long:               `Update remote refs along with associated objects`,
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		isAutoCommit := viper.GetBool("autoCommit")
		if isAutoCommit {
			options := append([]string{utils.GitDir, utils.WorkTree, "commit", "-am", `"updated by dotman"`}, []string{}...)
			utils.GitCommand(true, "", options...)
		}

		options := append([]string{utils.GitDir, utils.WorkTree, "push"}, args...)
		utils.GitCommand(true, "", options...)
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pushCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pushCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
