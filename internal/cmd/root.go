package cmd

import (
	"github.com/spf13/cobra"

	"github.com/umbrella-sh/simply-dns-sync/internal/api"
	"github.com/umbrella-sh/simply-dns-sync/internal/cmd/backup"
	"github.com/umbrella-sh/simply-dns-sync/internal/cmd/push"
	"github.com/umbrella-sh/simply-dns-sync/internal/configs"
)

var rootCmd = &cobra.Command{
	Use:   configs.AppName,
	Short: configs.AppNameTitle,
}

func init() {
	rootCmd.AddCommand(backup.BackupCmd)
	rootCmd.AddCommand(push.PushCmd)
}

func RootExecute() error {
	api.Init(api.SimplyApiConfig{
		Url:           configs.Main.SimplyApi.Url,
		AccountNumber: configs.Main.SimplyApi.AccountNumber,
		AccountApiKey: configs.Main.SimplyApi.AccountApiKey,
	})
	
	return rootCmd.Execute()
}
