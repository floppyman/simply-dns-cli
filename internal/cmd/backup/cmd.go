package backup

import "github.com/spf13/cobra"

//goland:noinspection GoNameStartsWithPackageName
var BackupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Pulls all the current Domains and DNS records from account and stores them locally.",
	Args:  handleArgs,
	Run:   cmdRun,
}

func handleArgs(cmd *cobra.Command, args []string) error {
	return nil
}
