package forms

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
	log "github.com/floppyman/um-common/logging/basic"

	gf "github.com/floppyman/simply-dns-cli/internal/forms/generic_fields"
)

func RunAcceptInput() (bool, bool) {
	model := gf.InitGenericBooleanModel(gf.GenericBooleanModelInput{
		HeaderText:   "Is this correct?",
		InitialValue: true,
		Mode:         gf.GbmYesNo,
	})
	p := tea.NewProgram(model)
	m, err := p.Run()
	if err != nil {
		log.Errorln("tea failed, ", err)
		os.Exit(1)
	}
	if m, ok := m.(gf.GenericBooleanModel); ok && !m.InputCancelled() {
		return false, m.ChosenYes()
	}
	return true, false
}
