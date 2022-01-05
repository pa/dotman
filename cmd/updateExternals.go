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
		// TODO: if files get deleted in upstream
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

				var isUpToDate bool = false
				var isCloned bool = false
				var currentCommitID []byte
				var deletedFiles []string
				var addedAndModifiedFiles []string

				if externalsPaths != nil {
					// clone externals repo
					if utils.IsGitRepoDir(repoPath) {
						// Get current commit id
						currentCommitID, _ = utils.GitCommand(false,
							repoPath,
							"rev-parse",
							"HEAD",
						)

						// Download only repo objects and refs
						utils.GitCommand(false,
							repoPath,
							"fetch",
						)

						isUpToDate = utils.IsRepoUpToDate(repoPath)
					} else {
						utils.GitCommand(true,
							"",
							"clone",
							gitUrl,
							repoPath,
						)
						isCloned = true
					}

					if !isUpToDate {
						// Merge fetched objects and refs
						utils.GitCommand(false,
							repoPath,
							"merge",
						)

						// Get files deleted between previous commit to latest commit
						deletedFiles = utils.GetDeletedOrModifiedFiles(repoPath, currentCommitID, true)
						if len(deletedFiles) > 0 {
							for _, file := range deletedFiles {
								os.Remove(utils.HomeDir + "/" + externalKey + "/" + file)
							}
						}

						// Get files added and modified files between previous commit until latest commit
						addedAndModifiedFiles = utils.GetDeletedOrModifiedFiles(repoPath, currentCommitID, false)

						for _, externalsPath := range externalsPaths.([]interface{}) {
							source := strings.Split(externalsPath.(string), " ")[0]
							target := strings.Split(externalsPath.(string), " ")[1]
							sourcePath := utils.ExternalsDir + "/" + gitRepoBaseName + "/" + source
							targetPath := utils.HomeDir + "/" + externalKey + "/" + target

							// list files
							files, err := ioutil.ReadDir(sourcePath)
							if err != nil {
								// Create target dir if not exists
								utils.CreateDir(path.Dir(targetPath))

								// Only copy updated files
								if utils.Contains(addedAndModifiedFiles, source) || isCloned {
									utils.CopyFile(sourcePath, targetPath)
								}
							} else {
								for _, file := range files {
									// Create target dir if not exists
									utils.CreateDir(targetPath)

									_, err := os.Stat(targetPath + "/" + file.Name())
									if err == nil {
										// Only copy updated files
										if utils.Contains(addedAndModifiedFiles, source+"/"+file.Name()) || isCloned {
											os.Remove(targetPath + "/" + file.Name())
											utils.CopyFile(sourcePath+"/"+file.Name(), targetPath+"/"+file.Name())
										}
									} else {
										// Only copy updated files
										if utils.Contains(addedAndModifiedFiles, source+"/"+file.Name()) || isCloned {
											utils.CopyFile(sourcePath+"/"+file.Name(), targetPath+"/"+file.Name())
										}
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
						gitCmdOutput, _ := utils.GitCommand(false,
							repoPath,
							"pull",
						)

						fmt.Print(gitRepoBaseName, " - ", string(gitCmdOutput))
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
