package utils

import (
	"errors"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

type promptParams struct {
	label    string
	errorMsg string
}

func PromptGetInput(params promptParams) string {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(params.errorMsg)
		}
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	prompt := promptui.Prompt{
		Label:     params.label,
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	return result
}
