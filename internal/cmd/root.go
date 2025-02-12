package cmd

import (
	"github.com/spf13/cobra"

	"github.com/umbrella-sh/simply-dns-sync/internal/cmd/pull"
	"github.com/umbrella-sh/simply-dns-sync/internal/cmd/push"
	"github.com/umbrella-sh/simply-dns-sync/internal/configs"
)

var rootCmd = &cobra.Command{
	Use:   configs.AppName,
	Short: configs.AppNameTitle,
}

func init() {
	rootCmd.AddCommand(pull.PullCmd)
	rootCmd.AddCommand(push.PushCmd)
}

func RootExecute() error {
	return rootCmd.Execute()
}
