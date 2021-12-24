package utils

const (
	DotmanDirName         = ".dotman"
	DotFilesBackupDirName = "dotfiles-backup"
	GitUserName           = "dotman"
)

var HomeDir string = GetHomeDir()
var DotmanDir string = HomeDir + "/" + DotmanDirName
var GitDir string = "--git-dir=" + DotmanDir
var WorkTree string = "--work-tree=" + HomeDir
var DotfileBackupDir = DotmanDir + "/" + DotFilesBackupDirName
var EnableInteractiveCommand bool = false
