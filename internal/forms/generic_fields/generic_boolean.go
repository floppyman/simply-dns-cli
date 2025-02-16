package generic_fields

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/umbrella-sh/um-common/ext"

	"github.com/umbrella-sh/simply-dns-cli/internal/styles"
)

type (
	GenericBooleanModel struct {
		headerText string
		cancelForm bool
		choices    []string
		selected   bool
		mode       BooleanMode
	}
	BooleanMode int
)

const (
	GbmYesNo BooleanMode = iota
	GbmTrueFalse
	GbmAcceptDecline
)

var (
	gbmYesNoChoices         = []string{"Yes", "No"}
	gbmTrueFalseChoices     = []string{"True", "False"}
	gbmAcceptDeclineChoices = []string{"Accept", "Decline"}
)

func InitGenericBooleanModel(headerText string, mode BooleanMode, initialValue bool) GenericBooleanModel {
	//goland:noinspection SpellCheckingInspection
	model := GenericBooleanModel{
		headerText: headerText,
		selected:   initialValue,
		mode:       mode,
	}
	switch mode {
	case GbmYesNo:
		model.choices = gbmYesNoChoices
	case GbmTrueFalse:
		model.choices = gbmTrueFalseChoices
	case GbmAcceptDecline:
		model.choices = gbmAcceptDeclineChoices
	}
	return model
}

func (m GenericBooleanModel) Init() tea.Cmd {
	return nil
}

func (m GenericBooleanModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case tea.KeyUp.String(), "k": // move selection up
			m.selected = true

		case tea.KeyDown.String(), "j": // move selection down
			m.selected = false

		case tea.KeyEnter.String(): // choose selected
			return m, tea.Quit

		case tea.KeyEsc.String(), tea.KeyCtrlC.String(): // cancel form
			m.cancelForm = true
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m GenericBooleanModel) View() string {
	return fmt.Sprintf(`%s
  %s
  %s
`,
		styles.Header(m.headerText),
		ext.Iif(m.selected, styles.Input(m.choices[0]), styles.Normal(m.choices[0])),
		ext.Iif(!m.selected, styles.Input(m.choices[1]), styles.Normal(m.choices[1])),
	)
}

func (m GenericBooleanModel) InputCancelled() bool {
	return m.cancelForm
}

func (m GenericBooleanModel) ChosenYes() bool {
	return m.selected
}
