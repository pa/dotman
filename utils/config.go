package utils

const (
	dotmanDirName         = ".dotman"
	dotFilesBackupDirName = "dotfiles-backup"
)

var homeDir string = GetHomeDir()
var dotmanDir string = homeDir + "/" + dotmanDirName
var gitDir string = "--git-dir=" + dotmanDir
var workTree string = "--work-tree=" + homeDir
var dotfileBackupDir = dotmanDir + "/" + dotFilesBackupDirName
