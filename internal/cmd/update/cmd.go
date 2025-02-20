package update

import "github.com/spf13/cobra"

//goland:noinspection GoNameStartsWithPackageName
var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an existing DNS record on a Domain",
	Args:  handleArgs,
	Run:   cmdRun,
}

type updateOptions struct {
	Domain string
}

var options = updateOptions{
	Domain: "",
}

func init() {
	UpdateCmd.Flags().StringVarP(&options.Domain, "domain", "d", "", "TLD name to remove the record from, ex: domain.com")
}

func handleArgs(cmd *cobra.Command, args []string) error {
	return nil
}
