package utils

const (
	DotmanDirName         = ".dotman"
	DotFilesBackupDirName = "dotfiles-backup"
	GitUserName           = "dotman"
	ExternalsDirName      = "externals"
)

var HomeDir string = GetHomeDir()
var DotmanDir string = HomeDir + "/" + DotmanDirName
var GitDir string = "--git-dir=" + DotmanDir
var WorkTree string = "--work-tree=" + HomeDir
var DotfileBackupDir string = DotmanDir + "/" + DotFilesBackupDirName
var EnableInteractiveCommand bool = false
var ExternalsDir string = DotmanDir + "/" + ExternalsDirName
