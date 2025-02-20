package forms

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	log "github.com/umbrella-sh/um-common/logging/basic"

	gf "github.com/umbrella-sh/simply-dns-cli/internal/forms/generic_fields"
)

func RunNameInput() (bool, string) {
	p := tea.NewProgram(gf.InitGenericInputModel(gf.GenericInputModelInput{
		HeaderText:      fmt.Sprintf("%-*s", longestHeader, "Name:"),
		PlaceHolderText: "Ex. sub-domain",
		ValueCharLimit:  128,
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
