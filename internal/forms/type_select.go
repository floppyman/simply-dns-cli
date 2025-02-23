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

var TypeSelectHeader = fmt.Sprintf("%-*s", longestHeader, "Entry Type:")

func RunTypeSelect(initialValue apio.DnsRecordType) (bool, apio.DnsRecordType) {
	choices := []string{
		string(apio.DnsRecTypeA),
		string(apio.DnsRecTypeAAAA),
		string(apio.DnsRecTypeALIAS),
		string(apio.DnsRecTypeCAA),
		string(apio.DnsRecTypeCNAME),
		string(apio.DnsRecTypeDNSKEY),
		string(apio.DnsRecTypeDS),
		string(apio.DnsRecTypeHTTPS),
		string(apio.DnsRecTypeLOC),
		string(apio.DnsRecTypeMX),
		string(apio.DnsRecTypeNS),
		string(apio.DnsRecTypeSSHFP),
		string(apio.DnsRecTypeTLSA),
		string(apio.DnsRecTypeTXT),
	}
	values := []any{
		apio.DnsRecTypeA,
		apio.DnsRecTypeAAAA,
		apio.DnsRecTypeALIAS,
		apio.DnsRecTypeCAA,
		apio.DnsRecTypeCNAME,
		apio.DnsRecTypeDNSKEY,
		apio.DnsRecTypeDS,
		apio.DnsRecTypeHTTPS,
		apio.DnsRecTypeLOC,
		apio.DnsRecTypeMX,
		apio.DnsRecTypeNS,
		apio.DnsRecTypeSSHFP,
		apio.DnsRecTypeTLSA,
		apio.DnsRecTypeTXT,
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
		return false, m.Values[m.SelectedIndex()].(apio.DnsRecordType)
	}
	return true, ""
}
