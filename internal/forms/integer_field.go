package forms

import "strconv"

type IntegerConverter func(val string) (bool, int)
type IntegerFieldValidator func(val string, isRequired bool, converter IntegerConverter) (ok bool, msg string, res int)

type IntegerField struct {
	key            string
	headerText     string
	isRequired     bool
	value          int
	valueValidator IntegerFieldValidator
}

func NewIntegerField(key string, header string, isRequired bool, validator IntegerFieldValidator) *IntegerField {
	return &IntegerField{
		key:            key,
		headerText:     header,
		isRequired:     isRequired,
		value:          0,
		valueValidator: validator,
	}
}

func (f *IntegerField) Render() {
	for {
		printHeader(f.headerText, f.isRequired)

		input := readInput()

		ok, msg, val := f.valueValidator(input, f.isRequired, f.convertStringToInt)

		if !ok {
			printError(msg)
			continue
		}

		f.value = val
		break
	}
}

func (f *IntegerField) GetValue() (key string, val any) {
	return f.key, f.value
}

func (f *IntegerField) SetValue(val any) {
	f.value = val.(int)
}

func (f *IntegerField) convertStringToInt(val string) (bool, int) {
	res, err := strconv.Atoi(val)
	if err != nil {
		return false, 0
	}
	return true, res
}
