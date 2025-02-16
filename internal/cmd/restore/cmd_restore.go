package restore

import (
	"github.com/spf13/cobra"

	"github.com/umbrella-sh/simply-dns-cli/internal/styles"
)

//goland:noinspection GoNameStartsWithPackageName
var RestoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Push all the DNS record changes to a Domain based on a selected backup taken previously",
	Args:  handleArgs,
	Run:   cmdRun,
}

func handleArgs(cmd *cobra.Command, args []string) error {
	return nil
}

func cmdRun(_ *cobra.Command, _ []string) {
	// currentProducts := shared.PullProducts()
	styles.WarnPrint("Not implemented yet")
}
