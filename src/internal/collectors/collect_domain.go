package collectors

import (
	"fmt"

	"github.com/floppyman/simply-dns-cli/internal/forms"
	"github.com/floppyman/simply-dns-cli/internal/objects"
	"github.com/floppyman/simply-dns-cli/internal/shared"
	"github.com/floppyman/simply-dns-cli/internal/styles"
)

func AcceptInfo() (cancelled bool, accepted bool) {
	cancelled, accepted = forms.RunAcceptInput()
	if cancelled {
		return cancelled, accepted
	}
	return cancelled, accepted
}

func CollectDomain(initialDomain string) (cancelled bool, domain string) {
	if initialDomain == "" {
		objNames := shared.PullProductNames()
		styles.Blank()

		cancelled, domain = forms.RunDomainSelect(objNames)
		if cancelled {
			return cancelled, ""
		}
	} else {
		domain = initialDomain
		shared.PrintValue(forms.DomainSelectHeader, domain)
	}

	return
}

func CollectDnsRecord(initialDnsRecord int64, domain string) (cancelled bool, record *objects.SimplyDnsRecord) {
	records := shared.PullDnsRecords(domain, "")

	if initialDnsRecord > 0 {
		cancelled = false
		for _, v := range records {
			if v.RecordId == initialDnsRecord {
				record = v
				break
			}
		}
		if record == nil {
			shared.PrintValue(forms.DnsRecordSelectHeader, "No Record found with provided Id")
			return true, nil
		}
		shared.PrintValue(forms.DnsRecordSelectHeader, fmt.Sprintf("%10d | %s", record.RecordId, record.Name))
		return false, record
	}

	if domain == "" {
		printDomainIsRequired()
		return true, nil
	}

	choices := make([]string, 0)
	values := make([]any, 0)
	for _, v := range records {
		choices = append(choices, fmt.Sprintf("%-10d | %s", v.RecordId, v.Name))
		values = append(values, v)
	}
	styles.Blank()

	cancelled, record = forms.RunDnsRecordSelect(choices, values)
	if cancelled {
		return cancelled, nil
	}

	return
}

func printDomainIsRequired() {
	styles.Warn("Domain is required to print dns records")
}
