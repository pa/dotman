package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/pa/dotman/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// updateExternalsCmd represents the updateExternals command
var updateExternalsCmd = &cobra.Command{
	Use:   "update-externals",
	Short: "Downloads and updates git externals like plugins, etc",
	Long:  `Downloads and updates git externals like plugins, etc`,
	Run: func(cmd *cobra.Command, args []string) {
		// issue with unmarshalling the yaml config https://github.com/spf13/viper/issues/338
		// type configPathAttributes struct {
		// 	url   string   `mapstructure:"url"`
		// 	paths []string `mapstructure:"paths"`
		// }

		// type Externals struct {
		// 	configPath map[string][]configPathAttributes
		// }

		// type Config struct {
		// 	AutoCommit bool                   `mapstructure:"autoCommit"`
		// 	Externals  map[string]interface{} `mapstructure:"externals"`
		// }

		// var config Config

		// err := viper.Unmarshal(&config)
		// if err != nil {
		// 	fmt.Print("unable to decode into struct", err)
		// }

		// Create externals dir under dotman directory
		utils.CreateDir(utils.ExternalsDir)

		externalsData := viper.GetStringMap("externals")

		for externalKey, externalValue := range externalsData {
			for _, external := range externalValue.([]interface{}) {
				// convert interface to map
				external := external.(map[interface{}]interface{})

				gitUrl := external["url"].(string)
				gitRepoBaseName := strings.Split(path.Base(gitUrl), ".")[0]
				repoPath := utils.ExternalsDir + "/" + gitRepoBaseName
				externalsPaths := external["paths"]

				var isUpToDate []byte

				if externalsPaths != nil {
					// clone externals repo
					if utils.IsGitRepoDir(repoPath) {
						isUpToDate, _ = utils.GitCommand(false,
							repoPath,
							"pull",
						)
					} else {
						utils.GitCommand(true,
							"",
							"clone",
							gitUrl,
							repoPath,
						)
					}

					isUpToDate := utils.RemoveRunes(string(isUpToDate))

					if isUpToDate != "Already up to date." {
						for _, externalsPath := range externalsPaths.([]interface{}) {
							sourcePath := utils.ExternalsDir + "/" + gitRepoBaseName + "/" + strings.Split(externalsPath.(string), " ")[0]
							targetPath := utils.HomeDir + "/" + externalKey + "/" + strings.Split(externalsPath.(string), " ")[1]

							// list files
							files, err := ioutil.ReadDir(sourcePath)
							if err != nil {
								// Create target dir if not exists
								utils.CreateDir(path.Dir(targetPath))

								utils.CopyFile(sourcePath, targetPath)
							} else {
								for _, file := range files {
									// Create target dir if not exists
									utils.CreateDir(targetPath)

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
						fmt.Print(gitRepoBaseName, " - Already up to date.\n")
					}

				} else {
					repoPath := utils.HomeDir + "/" + externalKey + "/" + gitRepoBaseName
					// clone externals repo
					if utils.IsGitRepoDir(repoPath) {
						isUpToDate, _ = utils.GitCommand(false,
							repoPath,
							"pull",
						)

						fmt.Print(gitRepoBaseName, " - ", string(isUpToDate))
					} else {
						utils.GitCommand(true,
							"",
							"clone",
							gitUrl,
							repoPath,
						)
					}
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
