package generic_fields

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/umbrella-sh/um-common/ext"

	"github.com/umbrella-sh/simply-dns-cli/internal/styles"
)

type (
	errMsg                 error
	GenericInputConverter  func(val string) (bool, any)
	GenericInputValidator  func(val string, isRequired bool, valueConverter GenericInputConverter) (ok bool, msg string)
	GenericInputModelInput struct {
		HeaderText      string
		PlaceHolderText string
		Width           int
		ValueCharLimit  int
		InitialValue    string
		IsRequired      bool
		InputValidator  GenericInputValidator
		InputConverter  GenericInputConverter
	}
	GenericInputModel struct {
		GenericInputModelInput
		textInput  textinput.Model
		err        error
		cancelForm bool
	}
)

func InitGenericInputModel(model GenericInputModelInput) GenericInputModel {
	ti := textinput.New()
	ti.Placeholder = model.PlaceHolderText
	ti.CharLimit = ext.Iif(model.ValueCharLimit > 0, model.ValueCharLimit, 1000)
	ti.Width = ext.Iif(model.Width > 0, model.Width, 30)
	ti.Prompt = ""
	ti.SetValue(model.InitialValue)
	ti.Focus()

	return GenericInputModel{
		GenericInputModelInput: model,
		textInput:              ti,
		err:                    nil,
	}
}

func (m GenericInputModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m GenericInputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			m.textInput.Blur()
			return m, tea.Quit
		case tea.KeyCtrlC, tea.KeyEsc:
			m.cancelForm = true
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m GenericInputModel) View() string {
	ok, msg := m.InputValidator(m.textInput.Value(), m.IsRequired, m.InputConverter)
	if !ok {
		return fmt.Sprintf(
			"%s %s\n%s",
			styles.Header(m.HeaderText),
			m.textInput.View(),
			styles.Error(msg),
		) + "\n"
	}

	return fmt.Sprintf(
		"%s %s",
		styles.Header(m.HeaderText),
		m.textInput.View(),
	) + "\n"
}

func (m GenericInputModel) InputCancelled() bool {
	return m.cancelForm
}

func (m GenericInputModel) GetValue() string {
	return m.textInput.Value()
}

func (m GenericInputModel) GetValueConverted() any {
	_, convertedOutput := m.InputConverter(m.textInput.Value())
	return convertedOutput
}
