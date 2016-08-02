package form

const (
	InputCheckbox = "input:checkbox"
)

type FieldBridge struct {
	_type     string
	multi     bool
	label     *string
	name      *string
	suffix    **string
	id        string
	value     *string
	values    *[]string
	inputAttr map[string]string
	checked   bool
	options   []string
	meta      map[string]string
}

func NewFieldBridge(
	_type string,
	multi bool,
	label *string,
	name *string,
	suffix **string,
	value *string,
	values *[]string,
) *FieldBridge {
	return &FieldBridge{
		_type:     _type,
		multi:     multi,
		label:     label,
		name:      name,
		suffix:    suffix,
		value:     value,
		values:    values,
		inputAttr: map[string]string{},
		meta:      map[string]string{},
	}
}

func (f *FieldBridge) Id() string {
	return f.id
}

func (f *FieldBridge) Label() string {
	return *f.label
}

func (f *FieldBridge) Name() string {
	if nil != *f.suffix {
		return *f.name + **f.suffix
	}

	return *f.name
}

func (f *FieldBridge) Value() string {
	return *f.value
}

func (f *FieldBridge) Values() []string {
	return *f.value
}

func (f *FieldBridge) InputAttr() map[string]string {
	delete(f.inputAttr, "id")
	delete(f.inputAttr, "class")
	delete(f.inputAttr, "selected")
	delete(f.inputAttr, "checked")
	delete(f.inputAttr, "name")

	return f.inputAttr
}

func (f *FieldBridge) SetInputAttr(name, value string) {
	f.inputAttr[name] = value
}

func (f *FieldBridge) SetChecked(checked bool) {
	f.checked = checked
}

func (f *FieldBridge) Checked() bool {
	return f.checked
}

func (f *FieldBridge) SetOptions(options []string) {
	f.options = options
}

func (f *FieldBridge) Options() []string {
	return f.options
}

func (f *FieldBridge) HasMeta(name string) bool {
	return "" != f.meta[name]
}

func (f *FieldBridge) Meta(name string) string {
	return f.meta[name]
}

func (f *FieldBridge) SetMeta(name, value string) {
	f.meta[name] = value
}
