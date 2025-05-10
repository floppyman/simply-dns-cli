package forms

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	log "github.com/floppyman/um-common/logging/basic"

	gf "github.com/floppyman/simply-dns-cli/internal/forms/generic_fields"
	"github.com/floppyman/simply-dns-cli/internal/objects"
	"github.com/floppyman/simply-dns-cli/internal/shared"
)

var TtlSelectHeader = fmt.Sprintf("%-*s", longestHeader, "TTL:")

func RunTtlSelect(initialValue objects.DnsRecordTTL) (bool, objects.DnsRecordTTL) {
	choices := []string{
		"10 Minutes",
		"1 Hour (recommended)",
		"6 Hours",
		"12 Hours",
		"24 Hours",
	}
	values := []any{
		objects.DnsRecTTLMin10,
		objects.DnsRecTTLHour1,
		objects.DnsRecTTLHours6,
		objects.DnsRecTTLHours12,
		objects.DnsRecTTLHours24,
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
		return false, m.Values[m.SelectedIndex()].(objects.DnsRecordTTL)
	}
	return true, objects.DnsRecTTLHour1
}
