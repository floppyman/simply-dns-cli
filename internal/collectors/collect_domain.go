package collectors

import (
	"fmt"
	"strconv"

	"github.com/umbrella-sh/simply-dns-cli/internal/forms"
	"github.com/umbrella-sh/simply-dns-cli/internal/shared"
	"github.com/umbrella-sh/simply-dns-cli/internal/styles"
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

func CollectDnsRecord(initialDnsRecord int64, domain string) (cancelled bool, recordId int64) {
	if initialDnsRecord > 0 {
		cancelled = false
		recordId = initialDnsRecord
		shared.PrintValue(forms.DnsRecordSelectHeader, strconv.FormatInt(recordId, 10))
		return
	}

	if domain == "" {
		printDomainIsRequired()
		return true, 0
	}

	records := shared.PullDnsRecords(domain, "")
	choices := make([]string, 0)
	values := make([]any, 0)
	for _, v := range records {
		choices = append(choices, fmt.Sprintf("%-10d | %s", v.RecordId, v.Name))
		values = append(values, v.RecordId)
	}
	styles.Blank()

	cancelled, recordId = forms.RunDnsRecordSelect(choices, values)
	if cancelled {
		return cancelled, 0
	}

	return
}

func printDomainIsRequired() {
	styles.Warn("Domain is required to print dns records")
}
