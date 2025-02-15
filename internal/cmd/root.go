package cmd

import (
	"github.com/spf13/cobra"

	"github.com/umbrella-sh/simply-dns-cli/internal/api"
	"github.com/umbrella-sh/simply-dns-cli/internal/cmd/add"
	"github.com/umbrella-sh/simply-dns-cli/internal/cmd/backup"
	"github.com/umbrella-sh/simply-dns-cli/internal/cmd/remove"
	"github.com/umbrella-sh/simply-dns-cli/internal/cmd/restore"
	"github.com/umbrella-sh/simply-dns-cli/internal/cmd/update"
	"github.com/umbrella-sh/simply-dns-cli/internal/configs"
)

var rootCmd = &cobra.Command{
	Use:   configs.AppName,
	Short: "",
}

func init() {
	rootCmd.AddCommand(add.AddCmd)
	rootCmd.AddCommand(update.UpdateCmd)
	rootCmd.AddCommand(remove.RemoveCmd)
	rootCmd.AddCommand(backup.BackupCmd)
	rootCmd.AddCommand(restore.RestoreCmd)
}

func RootExecute() error {
	api.Init(api.SimplyApiConfig{
		Url:           configs.Main.SimplyApi.Url,
		AccountNumber: configs.Main.SimplyApi.AccountNumber,
		AccountApiKey: configs.Main.SimplyApi.AccountApiKey,
	})
	
	return rootCmd.Execute()
}
