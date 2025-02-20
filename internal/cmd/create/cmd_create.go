package create

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/umbrella-sh/um-common/jsons"

	"github.com/umbrella-sh/simply-dns-cli/internal/api"
	"github.com/umbrella-sh/simply-dns-cli/internal/forms"
	"github.com/umbrella-sh/simply-dns-cli/internal/shared"
	"github.com/umbrella-sh/simply-dns-cli/internal/styles"
)

func cmdRun(_ *cobra.Command, _ []string) {
	styles.Println(styles.Info("Add new dns record"))
	styles.Blank()

	cancelled, domain, record := collectInfo()
	if cancelled {
		printCancelText()
		return
	}

	var accepted bool
	cancelled, accepted = acceptInfo()

	if !accepted {
		printNotAcceptedText()
		return
	}

	createRecord(domain, record)
}

//goland:noinspection GoNameStartsWithPackageName
func createRecord(domain string, record *api.SimplyDnsRecord) {
	styles.WaitPrint("Creating dns entry")

	_, err := api.CreateDnsRecord(domain, record)
	if err != nil {
		styles.FailPrint("Failed to create DNS Entry")
		styles.FailPrint("Error: %v", err)
		return
	}

	styles.SuccessPrint("DNS Entry created on %s", domain)
}

func acceptInfo() (cancelled bool, accepted bool) {
	styles.Blank()
	cancelled, accepted = forms.RunAcceptInput()
	if cancelled {
		return cancelled, accepted
	}
	styles.Blank()
	return cancelled, accepted
}

func collectInfo() (cancelled bool, domain string, record *api.SimplyDnsRecord) {
	record = &api.SimplyDnsRecord{}

	if options.Domain == "" {
		products := shared.PullProducts()
		var objNames = make([]string, 0)
		for _, product := range products {
			objNames = append(objNames, product.Object)
		}
		styles.Blank()

		cancelled, domain = forms.RunDomainSelect(objNames)
		if cancelled {
			return cancelled, "", nil
		}
	} else {
		domain = options.Domain
	}

	if options.Type == "" {
		cancelled, record.Type = forms.RunTypeSelect()
		if cancelled {
			return cancelled, "", nil
		}
	} else {
		record.Type = api.DnsRecordType(options.Type)
	}

	if options.TTL <= 0 {
		cancelled, record.TTL = forms.RunTtlSelect()
		if cancelled {
			return cancelled, "", nil
		}
	} else {
		record.TTL = api.DnsRecordTTL(options.TTL)
	}

	if options.Name == "" {
		var name string
		cancelled, name = forms.RunNameInput()
		if cancelled {
			return cancelled, "", nil
		}
		record.Name = fmt.Sprintf("%s.%s", name, domain)
	} else {
		record.Name = fmt.Sprintf("%s.%s", options.Name, domain)
	}

	if options.Data == "" {
		cancelled, record.Data = forms.RunDataInput()
		if cancelled {
			return cancelled, "", nil
		}
	} else {
		record.Data = options.Data
	}

	if record.Type == api.DnsRecTypeMX {
		if options.Priority <= 0 {
			cancelled, record.Priority = forms.RunPriorityInput()
			if cancelled {
				return cancelled, "", nil
			}
		} else {
			record.Priority = jsons.NewJsonInt32(int32(options.Priority))
		}
	} else {
		record.Priority = jsons.NullJsonInt32()
	}

	if options.Comment == NoCommentValue {
		cancelled, record.Comment = forms.RunCommentInput()
		if cancelled {
			return cancelled, "", nil
		}
	} else {
		record.Comment = options.Comment
	}

	return cancelled, domain, record
}

func printCancelText() { styles.Println(styles.Warn("\nCreate was cancelled\n")) }
func printNotAcceptedText() {
	styles.Println(styles.Warn("\nInformation is not accepted and no record was created\n"))
}
