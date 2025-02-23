package forms

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	log "github.com/umbrella-sh/um-common/logging/basic"

	apio "github.com/umbrella-sh/simply-dns-cli/internal/api_objects"
	gf "github.com/umbrella-sh/simply-dns-cli/internal/forms/generic_fields"
)

var DnsRecordSelectHeader = fmt.Sprintf("%-*s", longestHeader, "Dns Record:")

func RunDnsRecordSelect(choices []string, values []any) (bool, *apio.SimplyDnsRecord) {
	model := gf.InitGenericSelectModel(gf.GenericSelectModelInput{
		HeaderText:   DnsRecordSelectHeader,
		Choices:      choices,
		Values:       values,
		InitialValue: 0,
	})
	p := tea.NewProgram(model)
	m, err := p.Run()
	if err != nil {
		log.Errorln("tea failed, ", err)
		os.Exit(1)
	}
	if m, ok := m.(gf.GenericSelectModel); ok && !m.InputCancelled() {
		return false, m.Values[m.SelectedIndex()].(*apio.SimplyDnsRecord)
	}
	return true, nil
}
