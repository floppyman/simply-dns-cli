package remove

import (
	"github.com/spf13/cobra"

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

	cancelled, recordId := collectors.CollectDnsRecord(options.RecordId, domain)
	if cancelled {
		printCancelText()
		return
	}
	styles.Blank()

	styles.Blank()
	var accepted bool
	cancelled, accepted = collectors.AcceptInfo()
	if !accepted {
		printNotAcceptedText()
		return
	}
	styles.Blank()

	styles.Normal("%d", recordId)
}

func printCancelText() { styles.Println(styles.Warn("\nRemove was cancelled\n")) }
func printNotAcceptedText() {
	styles.Println(styles.Warn("\nInformation is not accepted and no record was created\n"))
}
