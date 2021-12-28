/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/pa/dotman/utils"
	"github.com/spf13/cobra"
)

// resetCmd represents the reset command
var resetCmd = &cobra.Command{
	Use:                "reset",
	Short:              "Native git reset command",
	Long:               `Native git reset command`,
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		options := append([]string{utils.GitDir, utils.WorkTree, "reset"}, args...)
		utils.GitCommandRun(options...)
	},
}

func init() {
	rootCmd.AddCommand(resetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// resetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// resetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
