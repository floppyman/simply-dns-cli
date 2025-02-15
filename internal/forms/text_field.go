package forms

type TextFieldValidator func(val string, isRequired bool) (ok bool, msg string, res string)

type TextField struct {
	key            string
	headerText     string
	isRequired     bool
	value          string
	valueValidator TextFieldValidator
}

func NewTextField(key string, header string, isRequired bool, validator TextFieldValidator) *TextField {
	return &TextField{
		key:            key,
		headerText:     header,
		isRequired:     isRequired,
		value:          "",
		valueValidator: validator,
	}
}

func (f *TextField) Render() {
	for {
		printHeader(f.headerText, f.isRequired)

		input := readInput()

		ok, msg, val := f.valueValidator(input, f.isRequired)
		if !ok {
			printError(msg)
			continue
		}

		f.value = val
		break
	}
}

func (f *TextField) GetValue() (key string, val any) {
	return f.key, f.value
}

func (f *TextField) SetValue(val any) {
	f.value = val.(string)
}
