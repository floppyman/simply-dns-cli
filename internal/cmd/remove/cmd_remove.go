package remove

import (
	"github.com/spf13/cobra"

	"github.com/umbrella-sh/simply-dns-cli/internal/styles"
)

//goland:noinspection GoNameStartsWithPackageName
var RemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove an existing DNS record from a Domain",
	Args:  handleArgs,
	Run:   cmdRun,
}

func handleArgs(cmd *cobra.Command, args []string) error {
	return nil
}

func cmdRun(_ *cobra.Command, _ []string) {
	styles.WarnPrint("Not implemented yet")
}
