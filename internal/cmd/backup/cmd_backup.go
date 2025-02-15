package backup

import (
	"time"

	"github.com/spf13/cobra"
	"github.com/umbrella-sh/um-common/logging/ulog"
	"github.com/umbrella-sh/um-common/utils"

	"github.com/umbrella-sh/simply-dns-cli/internal/api"
	"github.com/umbrella-sh/simply-dns-cli/internal/configs"
)

//goland:noinspection GoNameStartsWithPackageName
var BackupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Pulls all the current Products and DNS Records from Simply.com DNS for local backup",
	Args:  handleCommandArguments,
	Run:   backupCmdRun,
}

func handleCommandArguments(cmd *cobra.Command, args []string) error {
	return nil
}

func backupCmdRun(_ *cobra.Command, _ []string) {
	ulog.Console.Info().Msg("Pulling products from account...")
	products, err := api.GetProducts()
	if err != nil {
		ulog.Console.Err(err).Msg("Failed to get products")
		return
	}
	ulog.Console.Info().Msg(utils.Green("Done"))

	ulog.Console.Info().Msg("Pulling dns records from each product...")
	for _, product := range products {
		ulog.Console.Info().Msgf("  %s", product.Name)
		records, err := api.GetDnsRecords(product.Object)
		if err != nil {
			ulog.Console.Err(err).Msgf("  Failed to get dns records for %s", product.Name)
			product.DnsRecords = make([]*api.SimplyDnsRecord, 0)
			continue
		}

		product.DnsRecords = records
	}
	ulog.Console.Info().Msg(utils.Green("Done"))

	ulog.Console.Info().Msg("Saving products to backup file...")
	err = configs.SaveBackup(products, time.Now())
	if err != nil {
		return
	}
	ulog.Console.Info().Msg(utils.Green("Done"))
}
