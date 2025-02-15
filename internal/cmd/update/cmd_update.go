package update

import (
	"github.com/spf13/cobra"
	log "github.com/umbrella-sh/um-common/logging/basic"
)

//goland:noinspection GoNameStartsWithPackageName
var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an existing DNS record on a Domain",
	Args:  handleArgs,
	Run:   cmdRun,
}

func handleArgs(cmd *cobra.Command, args []string) error {
	return nil
}

func cmdRun(_ *cobra.Command, _ []string) {
	log.Warnln("Not implemented yet")
}
