package main

import (
	"strings"

	"github.com/floppyman/simply-dns-cli/internal/cmd"
	"github.com/floppyman/simply-dns-cli/internal/configs"
	"github.com/floppyman/simply-dns-cli/internal/styles"
)

var (
	Version   = "0.0.0"
	BuildDate = "now"
	Debug     = "false"
	IsDebug   = strings.ToLower(Debug) == "true"
)

func main() {
	printHeader()

	err := configs.InitConfig(IsDebug)
	if err != nil {
		styles.FailPrint("Be sure to create a 'config.json' either in '~/.config/%s/' or besides the executable", configs.AppName)
		return
	}
	if IsDebug {
		styles.Blank()
	}

	// https://github.com/spf13/cobra/blob/v1.8.0/site/content/user_guide.md
	_ = cmd.RootExecute()
	styles.Blank()
}

func printHeader() {
	styles.Blank()
	styles.Print(styles.ProgramTitle(configs.AppNameTitle))
	styles.Print(styles.GraphicLight(" v%s", Version))
	styles.Println(styles.Graphic(" @ %s", BuildDate))
	styles.Blank()
}
