package pull

import (
	"github.com/spf13/cobra"
	"github.com/umbrella-sh/um-common/logging/ulog"
	"github.com/umbrella-sh/um-common/utils"

	"github.com/umbrella-sh/simply-dns-sync/internal/api"
	"github.com/umbrella-sh/simply-dns-sync/internal/configs"
)

//goland:noinspection GoNameStartsWithPackageName
var PullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pulls all the current Products and DNS Records from Simply.com DNS",
	Args:  handleCommandArguments,
	Run:   pullCmdRun,
}

func handleCommandArguments(cmd *cobra.Command, args []string) error {
	return nil
}

func pullCmdRun(_ *cobra.Command, _ []string) {
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

	ulog.Console.Info().Msg("Saving products to sync-state...")
	configs.Sync.Products = products
	err = configs.SaveSyncState()
	if err != nil {
		return
	}
	ulog.Console.Info().Msg(utils.Green("Done"))
}
