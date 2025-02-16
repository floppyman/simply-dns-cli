package backup

import (
	"time"

	"github.com/spf13/cobra"

	"github.com/umbrella-sh/simply-dns-cli/internal/shared"
	"github.com/umbrella-sh/simply-dns-cli/internal/styles"
)

//goland:noinspection GoNameStartsWithPackageName
var BackupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Pulls all the current Domains and DNS records and stores them locally",
	Args:  handleArgs,
	Run:   cmdRun,
}

func handleArgs(cmd *cobra.Command, args []string) error {
	return nil
}

func cmdRun(_ *cobra.Command, _ []string) {
	products := shared.PullProductsAndDnsRecords()
	if products == nil {
		styles.FailPrint("failed to get products")
		return
	}

	styles.WaitPrint("saving to backup file")
	fileName, err := SaveBackup(products, time.Now())
	if err != nil {
		styles.FailPrint("failed to save backup")
		styles.FailPrint("error: %v", err)
		return
	}
	styles.SuccessPrint("backup file saved, name: %s", fileName)
}
