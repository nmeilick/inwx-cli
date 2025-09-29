package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ConfirmationResult int

const (
	ConfirmationYes ConfirmationResult = iota
	ConfirmationNo
	ConfirmationAll
	ConfirmationCancel
)

func AskConfirmation(message string, skipPrompt bool) (ConfirmationResult, error) {
	if skipPrompt {
		return ConfirmationYes, nil
	}

	fmt.Printf("%s (Y)es/(N)o/(A)ll/(C)ancel: ", message)

	reader := bufio.NewReader(os.Stdin)
	response, err := reader.ReadString('\n')
	if err != nil {
		return ConfirmationCancel, err
	}

	response = strings.TrimSpace(strings.ToLower(response))

	switch response {
	case "y", "yes":
		return ConfirmationYes, nil
	case "n", "no":
		return ConfirmationNo, nil
	case "a", "all":
		return ConfirmationAll, nil
	case "c", "cancel":
		return ConfirmationCancel, nil
	default:
		return ConfirmationCancel, fmt.Errorf("invalid response")
	}
}

func AskSimpleConfirmation(message string, skipPrompt bool) (bool, error) {
	if skipPrompt {
		return true, nil
	}

	fmt.Printf("%s (y/N): ", message)

	reader := bufio.NewReader(os.Stdin)
	response, err := reader.ReadString('\n')
	if err != nil {
		return false, err
	}

	response = strings.TrimSpace(strings.ToLower(response))
	return response == "y" || response == "yes", nil
}
