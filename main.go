package main

import (
	"github.com/umbrella-sh/simply-dns-cli/internal/cmd"
	"github.com/umbrella-sh/simply-dns-cli/internal/configs"
	"github.com/umbrella-sh/simply-dns-cli/internal/styles"
)

var (
	Version   = "0.0.0"
	BuildDate = "now"
)

func main() {
	printHeader()

	err := configs.InitConfig()
	if err != nil {
		styles.FailPrint("be sure to create a 'config.json' either in '~/.config/%s/' or besides the executable", configs.AppName)
		return
	}
	styles.Blank()

	// https://github.com/spf13/cobra/blob/v1.8.0/site/content/user_guide.md
	_ = cmd.RootExecute()
	styles.Blank()
}

func printHeader() {
	styles.Blank()
	styles.Print(styles.Info(configs.AppNameTitle))
	styles.Print(styles.GraphicLight(" v%s", Version))
	styles.Println(styles.Graphic(" @ %s", BuildDate))
	styles.Blank()
}
