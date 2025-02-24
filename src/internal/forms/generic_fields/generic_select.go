package generic_fields

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/umbrella-sh/simply-dns-cli/internal/styles"
)

type (
	GenericSelectModelInput struct {
		HeaderText   string
		Choices      []string
		Values       []any
		InitialValue int
	}

	GenericSelectModel struct {
		GenericSelectModelInput
		cancelForm   bool
		finishedForm bool
		selected     int
	}
)

func InitGenericSelectModel(model GenericSelectModelInput) GenericSelectModel {
	res := GenericSelectModel{
		GenericSelectModelInput: model,
		cancelForm:              false,
		finishedForm:            false,
		selected:                model.InitialValue,
	}

	if res.selected < 0 {
		res.selected = 0
	}
	if res.selected > len(model.Choices) {
		res.selected = len(model.Choices) - 1
	}

	return res
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
			m.finishedForm = true
			return m, tea.Quit

		case tea.KeyEsc.String(), tea.KeyCtrlC.String(): // cancel form
			m.cancelForm = true
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m GenericSelectModel) View() string {
	if m.finishedForm {
		s := fmt.Sprintf("%s %s", styles.Header(m.HeaderText), m.Choices[m.selected])
		s += "\n"
		return s
	}

	s := fmt.Sprintf("%s", styles.Header(m.HeaderText))

	for i, choice := range m.Choices {
		cursor := " "

		s += "\n"

		if m.selected == i {
			cursor = ">"
			s += fmt.Sprintf("%s", styles.Input(fmt.Sprintf("%s %s", cursor, choice)))
			continue
		}

		s += fmt.Sprintf("%s", styles.Normal(fmt.Sprintf("%s %s", cursor, choice)))
	}

	return s
}

func (m GenericSelectModel) InputCancelled() bool {
	return m.cancelForm
}

func (m GenericSelectModel) SelectedIndex() int {
	return m.selected
}
