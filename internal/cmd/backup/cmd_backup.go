package backup

import (
	"time"

	"github.com/spf13/cobra"
	log "github.com/umbrella-sh/um-common/logging/basic"

	"github.com/umbrella-sh/simply-dns-cli/internal/shared"
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
		log.FailPrint("failed to get products")
		return
	}

	log.WaitPrint("saving to backup file")
	fileName, err := SaveBackup(products, time.Now())
	if err != nil {
		log.FailPrint("failed to save backup")
		log.Errorln(err)
		return
	}
	log.SuccessPrintf("backup file saved, name: %s\n", fileName)
}
