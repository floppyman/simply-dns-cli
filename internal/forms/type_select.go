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

var TypeSelectHeader = fmt.Sprintf("%-*s", longestHeader, "Entry Type:")

func RunTypeSelect(initialValue api.DnsRecordType) (bool, api.DnsRecordType) {
	choices := []string{
		string(api.DnsRecTypeA),
		string(api.DnsRecTypeAAAA),
		string(api.DnsRecTypeALIAS),
		string(api.DnsRecTypeCAA),
		string(api.DnsRecTypeCNAME),
		string(api.DnsRecTypeDNSKEY),
		string(api.DnsRecTypeDS),
		string(api.DnsRecTypeHTTPS),
		string(api.DnsRecTypeLOC),
		string(api.DnsRecTypeMX),
		string(api.DnsRecTypeNS),
		string(api.DnsRecTypeSSHFP),
		string(api.DnsRecTypeTLSA),
		string(api.DnsRecTypeTXT),
	}
	values := []any{
		api.DnsRecTypeA,
		api.DnsRecTypeAAAA,
		api.DnsRecTypeALIAS,
		api.DnsRecTypeCAA,
		api.DnsRecTypeCNAME,
		api.DnsRecTypeDNSKEY,
		api.DnsRecTypeDS,
		api.DnsRecTypeHTTPS,
		api.DnsRecTypeLOC,
		api.DnsRecTypeMX,
		api.DnsRecTypeNS,
		api.DnsRecTypeSSHFP,
		api.DnsRecTypeTLSA,
		api.DnsRecTypeTXT,
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
		return false, m.Values[m.SelectedIndex()].(api.DnsRecordType)
	}
	return true, ""
}
