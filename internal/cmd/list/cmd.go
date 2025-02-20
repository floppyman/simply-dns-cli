package list

import "github.com/spf13/cobra"

//goland:noinspection GoNameStartsWithPackageName
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all records from a domain",
	Args:  handleArgs,
	Run:   cmdRun,
}

type listOptions struct {
	Domain string
}

var options = listOptions{
	Domain: "",
}

func init() {
	ListCmd.Flags().StringVarP(&options.Domain, "domain", "d", "", "TLD name to remove the record from, ex: domain.com")
}

func handleArgs(cmd *cobra.Command, args []string) error {
	return nil
}
