package forms

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	log "github.com/floppyman/um-common/logging/basic"

	gf "github.com/floppyman/simply-dns-cli/internal/forms/generic_fields"
)

var CommentInputHeader = fmt.Sprintf("%-*s", longestHeader, "Comment:")

func RunCommentInput(initialValue string) (bool, string) {
	p := tea.NewProgram(gf.InitGenericInputModel(gf.GenericInputModelInput{
		HeaderText:      CommentInputHeader,
		PlaceHolderText: "",
		ValueCharLimit:  255,
		InitialValue:    initialValue,
		IsRequired:      false,
		InputValidator:  validateCommentInput,
		InputConverter:  nil,
	}))
	m, err := p.Run()
	if err != nil {
		log.Errorln("tea failed, ", err)
		os.Exit(1)
	}
	if m, ok := m.(gf.GenericInputModel); ok && !m.InputCancelled() {
		return false, m.GetValue()
	}
	return true, ""
}

func validateCommentInput(val string, required bool, valueConverter gf.GenericInputConverter) (ok bool, msg string) {
	if !required && val == "" {
		return true, ""
	}

	if len(val) > 255 {
		return false, "Comment cannot be longer than 255 chars"
	}

	return true, ""
}
