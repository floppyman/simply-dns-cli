package styles

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

// https://www.ditig.com/publications/256-colors-cheat-sheet
const (
	mediumOrchid1Color = lipgloss.Color("207")
	hotPinkColor       = lipgloss.Color("205")
	gray93Color        = lipgloss.Color("255")
	red1Color          = lipgloss.Color("196")
	dodgerBlue1Color   = lipgloss.Color("33")
	grey42Color        = lipgloss.Color("242")
	grey78Color        = lipgloss.Color("251")
	yellow1Color       = lipgloss.Color("226")
	cyan1Color         = lipgloss.Color("51")
	green1Color        = lipgloss.Color("46")
	orange1Color       = lipgloss.Color("214")
)

var (
	styleHeader       = lipgloss.NewStyle().Foreground(mediumOrchid1Color).Render
	styleInput        = lipgloss.NewStyle().Foreground(hotPinkColor).Render
	styleNormal       = lipgloss.NewStyle().Foreground(gray93Color).Render
	styleRequired     = lipgloss.NewStyle().Foreground(red1Color).Render
	styleError        = lipgloss.NewStyle().Foreground(red1Color).Render
	styleValue        = lipgloss.NewStyle().Foreground(dodgerBlue1Color).Render
	styleGraphic      = lipgloss.NewStyle().Foreground(grey42Color).Render
	styleGraphicLight = lipgloss.NewStyle().Foreground(grey78Color).Render
	styleInfo         = lipgloss.NewStyle().Foreground(cyan1Color).Render
	styleWarn         = lipgloss.NewStyle().Foreground(yellow1Color).Render
	styleSuccess      = lipgloss.NewStyle().Foreground(green1Color).Render
	styleProgramTitle = lipgloss.NewStyle().Foreground(orange1Color).Render
)

func Print(text string)                      { fmt.Print(text) }
func Printf(format string, a ...interface{}) { fmt.Printf(format, a...) }
func Println(text string)                    { fmt.Println(text) }
func ProgramTitle(format string, a ...interface{}) string {
	return styleProgramTitle(fmt.Sprintf(format, a...))
}
func Header(format string, a ...interface{}) string { return styleHeader(fmt.Sprintf(format, a...)) }
func Input(format string, a ...interface{}) string  { return styleInput(fmt.Sprintf(format, a...)) }
func Normal(format string, a ...interface{}) string { return styleNormal(fmt.Sprintf(format, a...)) }
func Required(format string, a ...interface{}) string {
	return styleRequired(fmt.Sprintf(format, a...))
}
func Error(format string, a ...interface{}) string   { return styleError(fmt.Sprintf(format, a...)) }
func Value(format string, a ...interface{}) string   { return styleValue(fmt.Sprintf(format, a...)) }
func Graphic(format string, a ...interface{}) string { return styleGraphic(fmt.Sprintf(format, a...)) }
func GraphicLight(format string, a ...interface{}) string {
	return styleGraphicLight(fmt.Sprintf(format, a...))
}
func Info(format string, a ...interface{}) string    { return styleInfo(fmt.Sprintf(format, a...)) }
func Warn(format string, a ...interface{}) string    { return styleWarn(fmt.Sprintf(format, a...)) }
func Success(format string, a ...interface{}) string { return styleSuccess(fmt.Sprintf(format, a...)) }

func InfoPrint(format string, a ...interface{}) {
	Print(Graphic("["))
	Print(Info(" info "))
	Print(Graphic("] "))
	Println(Normal(format, a...))
}

func WarnPrint(format string, a ...interface{}) {
	Print(Graphic("["))
	Print(Warn(" warn "))
	Print(Graphic("] "))
	Println(Normal(format, a...))
}

func SuccessPrint(format string, a ...interface{}) {
	Print(Graphic("["))
	Print(Success("  ok  "))
	Print(Graphic("] "))
	Println(Normal(format, a...))
}

func BlankPrint(format string, a ...interface{}) {
	Print(Graphic("["))
	Print(Normal("      "))
	Print(Graphic("] "))
	Println(Normal(format, a...))
}

func FailPrint(format string, a ...interface{}) {
	Print(Graphic("["))
	Print(Error(" fail "))
	Print(Graphic("] "))
	Println(Normal(format, a...))
}

func WaitPrint(format string, a ...interface{}) {
	Print(Graphic("["))
	Print(Info(" wait "))
	Print(Graphic("] "))
	Println(Normal(format, a...))
}

func Blank() {
	fmt.Println()
}
