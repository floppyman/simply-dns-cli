package forms

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
	log "github.com/umbrella-sh/um-common/logging/basic"

	gf "github.com/umbrella-sh/simply-dns-cli/internal/forms/generic_fields"
)

func RunDataInput() (bool, string) {
	p := tea.NewProgram(gf.InitGenericInputModel(gf.GenericInputModelInput{
		HeaderText:      "Data:",
		PlaceHolderText: "Ex. 127.0.0.1",
		ValueCharLimit:  255,
		IsRequired:      true,
		InputValidator:  validateDataInput,
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

func validateDataInput(val string, required bool, valueConverter gf.GenericInputConverter) (ok bool, msg string) {
	if !required && val == "" {
		return true, ""
	}

	if val == "" {
		return false, "Data is required"
	}

	if len(val) > 255 {
		return false, "Data cannot be longer than 255 chars"
	}

	return true, ""
}
