package push

import (
	"github.com/spf13/cobra"
	"github.com/umbrella-sh/um-common/logging/ulog"
	"github.com/umbrella-sh/um-common/utils"

	"github.com/umbrella-sh/simply-dns-cli/internal/api"
)

//goland:noinspection GoNameStartsWithPackageName
var PushCmd = &cobra.Command{
	Use:   "push",
	Short: "Pushes all the DNS Record changes to Simply.com DNS",
	Args:  handleCommandArguments,
	Run:   pushCmdRun,
}

func handleCommandArguments(cmd *cobra.Command, args []string) error {
	return nil
}

func pushCmdRun(_ *cobra.Command, _ []string) {
	// currentState := pullCurrentState()

}

func pullCurrentState() []*api.SimplyProduct {
	ulog.Console.Info().Msg("Pulling products from account...")
	products, err := api.GetProducts()
	if err != nil {
		ulog.Console.Err(err).Msg("Failed to get products")
		return nil
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

	return products
}
