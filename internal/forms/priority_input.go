package forms

import (
	"fmt"
	"os"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/umbrella-sh/um-common/jsons"
	log "github.com/umbrella-sh/um-common/logging/basic"

	gf "github.com/umbrella-sh/simply-dns-cli/internal/forms/generic_fields"
)

var PriorityInputHeader = fmt.Sprintf("%-*s", longestHeader, "Priority:")

func RunPriorityInput() (bool, *jsons.JsonInt32) {
	p := tea.NewProgram(gf.InitGenericInputModel(gf.GenericInputModelInput{
		HeaderText:      PriorityInputHeader,
		PlaceHolderText: "Ex. 10",
		ValueCharLimit:  255,
		InitialValue:    "",
		IsRequired:      true,
		InputValidator:  validatePriorityInput,
		InputConverter:  convertPriorityInput,
	}))
	m, err := p.Run()
	if err != nil {
		log.Errorln("tea failed, ", err)
		os.Exit(1)
	}
	if m, ok := m.(gf.GenericInputModel); ok && !m.InputCancelled() {
		return false, jsons.NewJsonInt32(m.GetValueConverted().(int32))
	}
	return true, nil
}

func validatePriorityInput(val string, required bool, valueConverter gf.GenericInputConverter) (ok bool, msg string) {
	if !required && val == "" {
		return true, ""
	}

	converted, converterOutput := valueConverter(val)
	if !converted {
		return false, "Priority is not a valid integer"
	}
	convertedVal := converterOutput.(int32)

	if convertedVal < 1 {
		return false, "Priority must be greater than 1"
	}

	if convertedVal > 65535 {
		return false, "Priority must be lesser than 65535"
	}

	return true, ""
}

func convertPriorityInput(val string) (bool, any) {
	if val == "" {
		return true, int32(0)
	}
	res, err := strconv.Atoi(val)
	if err != nil {
		return false, int32(0)
	}
	return true, int32(res)
}
