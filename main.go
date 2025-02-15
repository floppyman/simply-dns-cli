package main

import (
	log "github.com/umbrella-sh/um-common/logging/basic"

	"github.com/umbrella-sh/simply-dns-cli/internal/cmd"
	"github.com/umbrella-sh/simply-dns-cli/internal/configs"
)

var (
	Version   = "0.0.0"
	BuildDate = "now"
)

func main() {
	printHeader()

	err := configs.InitConfig()
	if err != nil {
		log.Errorf("be sure to create a 'config.json' either in '~/.config/%s/' or besides the executable", configs.AppName)
		return
	}
	log.Println()

	// https://github.com/spf13/cobra/blob/v1.8.0/site/content/user_guide.md
	_ = cmd.RootExecute()
	log.Println()
}

func printHeader() {
	log.Println()
	log.Info(configs.AppNameTitle)
	log.Debugf(" v%s", Version)
	log.Tracef(" @ %s\n\n", BuildDate)
}
