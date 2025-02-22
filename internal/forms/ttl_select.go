package forms

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	log "github.com/umbrella-sh/um-common/logging/basic"

	"github.com/umbrella-sh/simply-dns-cli/internal/api"
	gf "github.com/umbrella-sh/simply-dns-cli/internal/forms/generic_fields"
	"github.com/umbrella-sh/simply-dns-cli/internal/shared"
)

var TtlSelectHeader = fmt.Sprintf("%-*s", longestHeader, "TTL:")

func RunTtlSelect(initialValue api.DnsRecordTTL) (bool, api.DnsRecordTTL) {
	choices := []string{
		"10 Minutes",
		"1 Hour (recommended)",
		"6 Hours",
		"12 Hours",
		"24 Hours",
	}
	values := []any{
		api.DnsRecTTLMin10,
		api.DnsRecTTLHour1,
		api.DnsRecTTLHours6,
		api.DnsRecTTLHours12,
		api.DnsRecTTLHours24,
	}
	model := gf.InitGenericSelectModel(gf.GenericSelectModelInput{
		HeaderText:   TtlSelectHeader,
		Choices:      choices,
		Values:       values,
		InitialValue: shared.Index(values, initialValue),
	})
	p := tea.NewProgram(model)
	m, err := p.Run()
	if err != nil {
		log.Errorln("tea failed, ", err)
		os.Exit(1)
	}
	if m, ok := m.(gf.GenericSelectModel); ok && !m.InputCancelled() {
		return false, m.Values[m.SelectedIndex()].(api.DnsRecordTTL)
	}
	return true, api.DnsRecTTLHour1
}
