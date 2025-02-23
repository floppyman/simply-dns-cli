package forms

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	log "github.com/umbrella-sh/um-common/logging/basic"

	gf "github.com/umbrella-sh/simply-dns-cli/internal/forms/generic_fields"
	"github.com/umbrella-sh/simply-dns-cli/internal/objects"
	"github.com/umbrella-sh/simply-dns-cli/internal/shared"
)

var TypeSelectHeader = fmt.Sprintf("%-*s", longestHeader, "Entry Type:")

func RunTypeSelect(initialValue objects.DnsRecordType) (bool, objects.DnsRecordType) {
	choices := []string{
		string(objects.DnsRecTypeA),
		string(objects.DnsRecTypeAAAA),
		string(objects.DnsRecTypeALIAS),
		string(objects.DnsRecTypeCAA),
		string(objects.DnsRecTypeCNAME),
		string(objects.DnsRecTypeDNSKEY),
		string(objects.DnsRecTypeDS),
		string(objects.DnsRecTypeHTTPS),
		string(objects.DnsRecTypeLOC),
		string(objects.DnsRecTypeMX),
		string(objects.DnsRecTypeNS),
		string(objects.DnsRecTypeSSHFP),
		string(objects.DnsRecTypeTLSA),
		string(objects.DnsRecTypeTXT),
	}
	values := []any{
		objects.DnsRecTypeA,
		objects.DnsRecTypeAAAA,
		objects.DnsRecTypeALIAS,
		objects.DnsRecTypeCAA,
		objects.DnsRecTypeCNAME,
		objects.DnsRecTypeDNSKEY,
		objects.DnsRecTypeDS,
		objects.DnsRecTypeHTTPS,
		objects.DnsRecTypeLOC,
		objects.DnsRecTypeMX,
		objects.DnsRecTypeNS,
		objects.DnsRecTypeSSHFP,
		objects.DnsRecTypeTLSA,
		objects.DnsRecTypeTXT,
	}
	model := gf.InitGenericSelectModel(gf.GenericSelectModelInput{
		HeaderText:   TypeSelectHeader,
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
		return false, m.Values[m.SelectedIndex()].(objects.DnsRecordType)
	}
	return true, ""
}
