package forms

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
	log "github.com/umbrella-sh/um-common/logging/basic"

	gf "github.com/umbrella-sh/simply-dns-cli/internal/forms/generic_fields"
)

func RunDomainSelect(choices []string) (bool, string) {
	values := make([]any, 0)
	for _, choice := range choices {
		values = append(values, choice)
	}
	p := tea.NewProgram(gf.InitGenericSelectModelWithDefault("Select Domain:", 0, choices, values))
	m, err := p.Run()
	if err != nil {
		log.Errorln("tea failed, ", err)
		os.Exit(1)
	}
	if m, ok := m.(gf.GenericSelectModel); ok && !m.InputCancelled() {
		return false, m.Values[m.SelectedIndex()].(string)
	}
	return true, ""
}
