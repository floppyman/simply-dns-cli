package generic_fields

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/umbrella-sh/simply-dns-cli/internal/styles"
)

type GenericSelectModel struct {
	headerText string
	cancelForm bool
	Choices    []string
	Values     []any
	selected   int
}

func InitGenericSelectModel(headerText string, choices []string, values []any) GenericSelectModel {
	//goland:noinspection SpellCheckingInspection
	return GenericSelectModel{
		headerText: headerText,
		Choices:    choices,
		Values:     values,
	}
}
func InitGenericSelectModelWithDefault(headerText string, defSelected int, choices []string, values []any) GenericSelectModel {
	//goland:noinspection SpellCheckingInspection
	return GenericSelectModel{
		headerText: headerText,
		Choices:    choices,
		Values:     values,
		selected:   defSelected,
	}
}

func (m GenericSelectModel) Init() tea.Cmd {
	return nil
}

func (m GenericSelectModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case tea.KeyUp.String(), "k": // move selection up
			if m.selected > 0 {
				m.selected--
			}

		case tea.KeyDown.String(), "j": // move selection down
			if m.selected < len(m.Choices)-1 {
				m.selected++
			}

		case tea.KeyEnter.String(): // choose selected
			return m, tea.Quit

		case tea.KeyEsc.String(), tea.KeyCtrlC.String(): // cancel form
			m.cancelForm = true
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m GenericSelectModel) View() string {
	s := fmt.Sprintf("%s\n", styles.Header(m.headerText))

	for i, choice := range m.Choices {
		cursor := " "
		if m.selected == i {
			cursor = ">"
			s += fmt.Sprintf("%s\n", styles.Input(fmt.Sprintf("%s %s", cursor, choice)))
			continue
		}

		s += fmt.Sprintf("%s\n", styles.Normal(fmt.Sprintf("%s %s", cursor, choice)))
	}

	return s
}

func (m GenericSelectModel) InputCancelled() bool {
	return m.cancelForm
}

func (m GenericSelectModel) SelectedIndex() int {
	return m.selected
}
