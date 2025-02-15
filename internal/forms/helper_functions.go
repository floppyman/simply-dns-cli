package forms

import (
	"fmt"

	log "github.com/umbrella-sh/um-common/logging/basic"
)

func printHeader(headerText string, isRequired bool) {
	log.Debug(headerText)
	if isRequired {
		log.Error(" *")
	}
	log.Print(": ")
}

func printError(text string) {
	log.Errorln(text)
}

func readInput() string {
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		input = ""
	}
	return input
}
