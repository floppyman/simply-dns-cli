package forms

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
	log "github.com/umbrella-sh/um-common/logging/basic"

	gf "github.com/umbrella-sh/simply-dns-cli/internal/forms/generic_fields"
)

func RunAcceptInput() (bool, bool) {
	p := tea.NewProgram(gf.InitGenericBooleanModel("Is this correct?", gf.GbmYesNo, true))
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
