package update

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/umbrella-sh/um-common/jsons"

	"github.com/umbrella-sh/simply-dns-cli/internal/api"
	"github.com/umbrella-sh/simply-dns-cli/internal/collectors"
	"github.com/umbrella-sh/simply-dns-cli/internal/forms"
	"github.com/umbrella-sh/simply-dns-cli/internal/shared"
	"github.com/umbrella-sh/simply-dns-cli/internal/styles"
)

func cmdRun(_ *cobra.Command, _ []string) {
	styles.Println(styles.Info("Update existing dns record"))
	styles.Blank()

	cancelled, domain, recordId, record := collectInfo()
	if cancelled {
		printCancelText()
		return
	}

	styles.Blank()
	var accepted bool
	cancelled, accepted = collectors.AcceptInfo()
	if !accepted {
		printNotAcceptedText()
		return
	}
	styles.Blank()

	updateRecord(domain, recordId, record)
}

//goland:noinspection GoNameStartsWithPackageName
func updateRecord(domain string, recordId int64, record *api.SimplyDnsRecord) {
	styles.WaitPrint("Updating dns entry")

	res, err := api.UpdateDnsRecord(domain, recordId, record)
	if err != nil {
		styles.FailPrint("Failed to update DNS Entry")
		styles.FailPrint("Error: %v", err)
		return
	}

	if res.Status != 200 {
		styles.FailPrint("Failed to update DNS Entry")
		styles.FailPrint("Error: %d, %v", res.Status, res.Message)
		return
	}

	styles.SuccessPrint("DNS Entry updated on %s", domain)
}

func collectInfo() (cancelled bool, domain string, recordId int64, record *api.SimplyDnsRecord) {
	record = &api.SimplyDnsRecord{}

	cancelled, domain = collectors.CollectDomain(options.Domain)
	if cancelled {
		return cancelled, "", 0, nil
	}
	styles.Blank()

	cancelled, record = collectors.CollectDnsRecord(options.RecordId, domain)
	if cancelled {
		return cancelled, "", 0, nil
	}
	recordId = record.RecordId
	record.RecordId = 0

	if options.Type == "" {
		cancelled, record.Type = forms.RunTypeSelect(record.Type)
		if cancelled {
			return cancelled, "", 0, nil
		}
	} else {
		record.Type = api.DnsRecordType(options.Type)
		shared.PrintValue(forms.TypeSelectHeader, api.DnsTypeToText(record.Type))
	}

	if options.TTL <= 0 {
		cancelled, record.TTL = forms.RunTtlSelect(record.TTL)
		if cancelled {
			return cancelled, "", 0, nil
		}
	} else {
		record.TTL = api.DnsRecordTTL(options.TTL)
		shared.PrintValue(forms.TtlSelectHeader, api.DnsTTLToText(record.TTL))
	}

	if options.Name == "" {
		var name string
		cancelled, name = forms.RunNameInput(record.Name)
		if cancelled {
			return cancelled, "", 0, nil
		}
		record.Name = fmt.Sprintf("%s.%s", name, domain)
	} else {
		record.Name = fmt.Sprintf("%s.%s", options.Name, domain)
		shared.PrintValue(forms.NameInputHeader, record.Name)
	}

	if options.Data == "" {
		cancelled, record.Data = forms.RunDataInput(record.Data)
		if cancelled {
			return cancelled, "", 0, nil
		}
	} else {
		record.Data = options.Data
		shared.PrintValue(forms.DataInputHeader, record.Data)
	}

	if record.Type == api.DnsRecTypeMX {
		if options.Priority <= 0 {
			cancelled, record.Priority = forms.RunPriorityInput(record.Priority)
			if cancelled {
				return cancelled, "", 0, nil
			}
		} else {
			record.Priority = jsons.NewJsonInt32(int32(options.Priority))
			shared.PrintValue(forms.PriorityInputHeader, record.Priority.ToString())
		}
	} else {
		record.Priority = jsons.NullJsonInt32()
	}

	if options.Comment == NoCommentValue {
		cancelled, record.Comment = forms.RunCommentInput(record.Comment)
		if cancelled {
			return cancelled, "", 0, nil
		}
	} else {
		record.Comment = options.Comment
		shared.PrintValue(forms.CommentInputHeader, record.Comment)
	}

	return
}

func printCancelText() { styles.Println(styles.Warn("\nCreate was cancelled")) }
func printNotAcceptedText() {
	styles.Println(styles.Warn("\nInformation is not accepted and no record was created"))
}
