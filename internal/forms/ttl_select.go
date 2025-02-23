package forms

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	log "github.com/umbrella-sh/um-common/logging/basic"

	apio "github.com/umbrella-sh/simply-dns-cli/internal/api_objects"
	gf "github.com/umbrella-sh/simply-dns-cli/internal/forms/generic_fields"
	"github.com/umbrella-sh/simply-dns-cli/internal/shared"
)

var TtlSelectHeader = fmt.Sprintf("%-*s", longestHeader, "TTL:")

func RunTtlSelect(initialValue apio.DnsRecordTTL) (bool, apio.DnsRecordTTL) {
	choices := []string{
		"10 Minutes",
		"1 Hour (recommended)",
		"6 Hours",
		"12 Hours",
		"24 Hours",
	}
	values := []any{
		apio.DnsRecTTLMin10,
		apio.DnsRecTTLHour1,
		apio.DnsRecTTLHours6,
		apio.DnsRecTTLHours12,
		apio.DnsRecTTLHours24,
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
		return false, m.Values[m.SelectedIndex()].(apio.DnsRecordTTL)
	}
	return true, apio.DnsRecTTLHour1
}
