package remove

import "github.com/spf13/cobra"

//goland:noinspection GoNameStartsWithPackageName
var RemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove an existing DNS record from a Domain",
	Args:  handleArgs,
	Run:   cmdRun,
}

type removeOptions struct {
	Domain   string
	RecordId int64
}

var options = removeOptions{
	Domain:   "",
	RecordId: 0,
}

func init() {
	RemoveCmd.Flags().StringVarP(&options.Domain, "domain", "d", "", "TLD name to remove the record from, ex: domain.com")
	RemoveCmd.Flags().Int64VarP(&options.RecordId, "record", "r", 0, "Id of the record to remove")
}

func handleArgs(cmd *cobra.Command, args []string) error {
	return nil
}
