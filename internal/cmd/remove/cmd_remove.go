package remove

import (
	"github.com/spf13/cobra"

	"github.com/umbrella-sh/simply-dns-cli/internal/api"
	"github.com/umbrella-sh/simply-dns-cli/internal/collectors"
	"github.com/umbrella-sh/simply-dns-cli/internal/styles"
)

func cmdRun(_ *cobra.Command, _ []string) {
	styles.Println(styles.Info("Remove existing dns record"))
	styles.Blank()

	cancelled, domain := collectors.CollectDomain(options.Domain)
	if cancelled {
		printCancelText()
		return
	}
	styles.Blank()

	cancelled, record := collectors.CollectDnsRecord(options.RecordId, domain)
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

	removeRecord(domain, record.RecordId)
}

//goland:noinspection GoNameStartsWithPackageName
func removeRecord(domain string, recordId int64) {
	styles.WaitPrint("Removing dns entry")

	res, err := api.DeleteDnsRecord(domain, recordId)
	if err != nil {
		styles.FailPrint("Failed to remove DNS Entry")
		styles.FailPrint("Error: %v", err)
		return
	}

	if res.Status != 200 {
		styles.FailPrint("Failed to remove DNS Entry")
		styles.FailPrint("Error: %d, %v", res.Status, res.Message)
		return
	}

	styles.SuccessPrint("DNS Entry removed on %s", domain)
}

func printCancelText() { styles.Println(styles.Warn("\nRemove was cancelled\n")) }
func printNotAcceptedText() {
	styles.Println(styles.Warn("\nInformation is not accepted and no record was created\n"))
}
