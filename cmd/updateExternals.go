package cmd

import (
	"io/ioutil"
	"os"
	"path"
	"reflect"
	"strings"

	"github.com/pa/dotman/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// updateExternalsCmd represents the updateExternals command
var updateExternalsCmd = &cobra.Command{
	Use:   "update-externals",
	Short: "Downloads and updates git externals like plugins, dependencies...",
	Long:  `Downloads and updates git externals like plugins, dependencies...`,
	Run: func(cmd *cobra.Command, args []string) {

		// Create externals dir under dotman directory
		utils.CreateDir(utils.ExternalsDir)

		externalsData := reflect.ValueOf(viper.GetStringMap("externals")).MapRange()

		for externalsData.Next() {
			gitUrl := viper.GetString("externals." + externalsData.Key().String() + ".url")
			externalsPaths := viper.GetStringSlice("externals." + externalsData.Key().String() + ".paths")
			gitRepoBaseName := strings.Split(path.Base(gitUrl), ".")[0]
			repoPath := utils.ExternalsDir + "/" + gitRepoBaseName

			if externalsPaths != nil {
				// clone externals repo
				if utils.IsGitRepoDir(repoPath) {
					utils.GitCommandRun(
						"pull",
					)
				} else {
					utils.GitCommandRun(
						"clone",
						gitUrl,
						repoPath,
					)
				}

				for _, externalsPath := range externalsPaths {
					sourcePath := utils.ExternalsDir + "/" + gitRepoBaseName + "/" + strings.Split(externalsPath, " ")[0]
					targetPath := utils.HomeDir + "/" + externalsData.Key().String() + "/" + strings.Split(externalsPath, " ")[1]

					// list files
					files, err := ioutil.ReadDir(sourcePath)
					if err != nil {
						// Create target dir if not exists
						_, dirErr := os.Stat(path.Dir(targetPath))
						if dirErr != nil {
							os.MkdirAll(path.Dir(targetPath), os.ModePerm)
						}
						utils.CopyFile(sourcePath, targetPath)
					} else {
						for _, file := range files {
							// Create target dir if not exists
							_, dirErr := os.Stat(targetPath)
							if dirErr != nil {
								os.MkdirAll(targetPath, os.ModePerm)
							}

							_, err := os.Stat(targetPath + "/" + file.Name())
							if err == nil {
								os.Remove(targetPath + "/" + file.Name())
								utils.CopyFile(sourcePath+"/"+file.Name(), targetPath+"/"+file.Name())
							} else {
								utils.CopyFile(sourcePath+"/"+file.Name(), targetPath+"/"+file.Name())

							}

						}
					}
				}
			} else {
				repoPath := utils.HomeDir + "/" + externalsData.Key().String() + "/" + gitRepoBaseName
				// clone externals repo
				if utils.IsGitRepoDir(repoPath) {
					utils.GitCommandRun(
						"pull",
					)
				} else {
					utils.GitCommandRun(
						"clone",
						gitUrl,
						utils.HomeDir+"/"+externalsData.Key().String()+"/"+gitRepoBaseName,
					)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(updateExternalsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateExternalsCmd.PersistentFlags().String("update-externals", "-ue", "Downloads and updates git externals in the dotman config")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateExternalsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
