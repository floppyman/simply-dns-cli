package forms

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	log "github.com/floppyman/um-common/logging/basic"

	gf "github.com/floppyman/simply-dns-cli/internal/forms/generic_fields"
)

var NameInputHeader = fmt.Sprintf("%-*s", longestHeader, "Name:")

func RunNameInput(initialValue string) (bool, string) {
	p := tea.NewProgram(gf.InitGenericInputModel(gf.GenericInputModelInput{
		HeaderText:      NameInputHeader,
		PlaceHolderText: "Ex. sub-domain",
		ValueCharLimit:  128,
		InitialValue:    initialValue,
		IsRequired:      false,
		InputValidator:  validateNameInput,
		InputConverter:  nil,
	}))
	m, err := p.Run()
	if err != nil {
		log.Errorln("tea failed, ", err)
		os.Exit(1)
	}
	if m, ok := m.(gf.GenericInputModel); ok && !m.InputCancelled() {
		return false, m.GetValue()
	}
	return true, ""
}

func validateNameInput(val string, required bool, valueConverter gf.GenericInputConverter) (ok bool, msg string) {
	if !required && val == "" {
		return true, ""
	}

	if val == "" {
		return false, "Name is required"
	}

	if len(val) > 128 {
		return false, "Name cannot be longer than 128 chars"
	}

	return true, ""
}
