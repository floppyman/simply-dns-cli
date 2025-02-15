package add

import (
	"github.com/spf13/cobra"
	log "github.com/umbrella-sh/um-common/logging/basic"
)

//goland:noinspection GoNameStartsWithPackageName
var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new DNS record to a Domain",
	Args:  handleArgs,
	Run:   cmdRun,
}

func handleArgs(cmd *cobra.Command, args []string) error {
	return nil
}

func cmdRun(_ *cobra.Command, _ []string) {
	log.Warnln("Not implemented yet")
}
