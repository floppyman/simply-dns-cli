package main

import (
	"github.com/umbrella-sh/um-common/logging/ulog"

	"github.com/umbrella-sh/simply-dns-sync/internal/cmd"
	"github.com/umbrella-sh/simply-dns-sync/internal/configs"
)

var (
	Version   = "0.0.0"
	BuildDate = "now"
)

func main() {
	ulog.New(100, 2, false)
	ulog.Console.Info().Msgf("%s v%s @ %s", configs.AppNameTitle, Version, BuildDate)
	ulog.Console.Info().Msg("")

	err := configs.InitConfig()
	if err != nil {
		ulog.Console.Error().Msgf("Be sure to create a 'config.json' either in '~/.config/%s/' or besides the executable", configs.AppName)
		return
	}
	err = configs.InitSync()
	if err != nil {
		return
	}

	// https://github.com/spf13/cobra/blob/v1.8.0/site/content/user_guide.md
	_ = cmd.RootExecute()
}
