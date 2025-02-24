package restore

import "github.com/spf13/cobra"

//goland:noinspection GoNameStartsWithPackageName
var RestoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Generate create or delete commands for each record that is different between the backup and api.",
	Args:  handleArgs,
	Run:   cmdRun,
}

type restoreOptions struct {
	Domain         string
	BackupFilePath string
}

var options = restoreOptions{
	Domain:         "",
	BackupFilePath: "",
}

func init() {
	RestoreCmd.Flags().StringVarP(&options.Domain, "domain", "d", "", "TLD name to remove the record from, ex: domain.com")
	RestoreCmd.Flags().StringVarP(&options.BackupFilePath, "backup-file-path", "f", "", "Name and path to the backup file to restore")
}

func handleArgs(cmd *cobra.Command, args []string) error {
	return nil
}
