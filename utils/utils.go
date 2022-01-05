package utils

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"
	"unicode"
)

func GetHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		fmt.Println(err)
	}
	return usr.HomeDir
}

func CreateDir(path string) error {
	var err error
	if _, err = os.Stat(path); os.IsNotExist(err) {
		if err = os.MkdirAll(path, os.ModePerm); err != nil {
			return err
		}
	}
	return err
}

func CopyFile(src, dst string) (err error) {
	r, err := os.Open(src)
	if err != nil {
		return err
	}
	defer r.Close()
	w, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer w.Close()
	w.ReadFrom(r)
	return nil
}

type NotInstalled struct {
	message string
	error
}

func (e *NotInstalled) Error() string {
	return e.message
}

func RunCommand(binary string, args ...string) (*exec.Cmd, error) {
	cmdExe, err := exec.LookPath(binary)
	if err != nil {
		if errors.Is(err, exec.ErrNotFound) {
			return nil, &NotInstalled{
				message: fmt.Sprintf("unable to find executable in PATH; please install %s before retrying", binary),
				error:   err,
			}
		}
		return nil, err
	}
	return exec.Command(cmdExe, args...), nil
}

func RemoveRunes(input string) string {
	cleanString := strings.Map(func(r rune) rune {
		if unicode.IsPrint(r) {
			return r
		}
		return -1
	}, input)
	return cleanString
}

// contains checks if a string is present in a slice
func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
