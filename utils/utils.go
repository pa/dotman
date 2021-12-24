package utils

import (
	"fmt"
	"os/user"
)

func GetHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		fmt.Println(err)
	}
	return usr.HomeDir
}
