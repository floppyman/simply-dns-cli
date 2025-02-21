package forms

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	log "github.com/umbrella-sh/um-common/logging/basic"

	"github.com/umbrella-sh/simply-dns-cli/internal/api"
	gf "github.com/umbrella-sh/simply-dns-cli/internal/forms/generic_fields"
)

var TtlSelectHeader = fmt.Sprintf("%-*s", longestHeader, "TTL:")

func RunTtlSelect() (bool, api.DnsRecordTTL) {
	choices := []string{
		"10 Minutes",
		"1 Hour (recommended)",
		"6 Hours",
		"12 Hours",
		"24 Hours",
	}
	values := []any{
		int(api.DnsRecTTLMin10),
		int(api.DnsRecTTLHour1),
		int(api.DnsRecTTLHours6),
		int(api.DnsRecTTLHours12),
		int(api.DnsRecTTLHours24),
	}
	model := gf.InitGenericSelectModel(gf.GenericSelectModelInput{
		HeaderText:   TtlSelectHeader,
		Choices:      choices,
		Values:       values,
		InitialValue: 1,
	})
	p := tea.NewProgram(model)
	m, err := p.Run()
	if err != nil {
		log.Errorln("tea failed, ", err)
		os.Exit(1)
	}
	if m, ok := m.(gf.GenericSelectModel); ok && !m.InputCancelled() {
		return false, api.DnsRecordTTL(m.Values[m.SelectedIndex()].(int))
	}
	return true, api.DnsRecTTLHour1
}
