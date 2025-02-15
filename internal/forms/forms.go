package forms

type IForm interface {
	Add(field IField, val any)
	Render()
	GetValues() map[string]any
}

type IField interface {
	Render()
	GetValue() (key string, val any)
	SetValue(val any)
}

type Form struct {
	IForm
	fields []IField
}

func New() Form {
	return Form{
		fields: make([]IField, 0),
	}
}

func (f *Form) Add(field IField, val any) {
	f.fields = append(f.fields, field)
	field.SetValue(val)
}

func (f *Form) Render() {
	for _, field := range f.fields {
		field.Render()
	}
}

func (f *Form) GetValues() map[string]any {
	res := make(map[string]any)
	for _, field := range f.fields {
		key, val := field.GetValue()
		res[key] = val
	}
	return res
}
