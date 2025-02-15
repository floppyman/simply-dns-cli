package restore

import (
	"github.com/spf13/cobra"
	log "github.com/umbrella-sh/um-common/logging/basic"
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
	log.Warnln("Not implemented yet")
}
