package form

type FieldAbstract struct {
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
}

func NewFieldAbstract(
	_type string,
	multi bool,
	label *string,
	name *string,
	suffix **string,
	value *string,
	values *[]string,
) *FieldAbstract {
	return &FieldAbstract{
		_type:     _type,
		multi:     multi,
		label:     label,
		name:      name,
		suffix:    suffix,
		value:     value,
		values:    values,
		inputAttr: map[string]string{},
	}
}

func (f *FieldAbstract) Id() string {
	return f.id
}

func (f *FieldAbstract) Label() string {
	return *f.label
}

func (f *FieldAbstract) Name() string {
	if nil != *f.suffix {
		return *f.name + **f.suffix
	}

	return *f.name
}

func (f *FieldAbstract) Value() string {
	return *f.value
}

func (f *FieldAbstract) Values() []string {
	return *f.value
}

func (f *FieldAbstract) InputAttr() map[string]string {
	delete(f.inputAttr, "id")
	delete(f.inputAttr, "class")
	delete(f.inputAttr, "selected")
	delete(f.inputAttr, "checked")
	delete(f.inputAttr, "name")

	return f.inputAttr
}

func (f *FieldAbstract) SetInputAttr(name, value string) {
	f.inputAttr[name] = value
}

func (f *FieldAbstract) SetChecked(checked bool) {
	f.checked = checked
}

func (f *FieldAbstract) Checked() bool {
	return f.checked
}

func (f *FieldAbstract) SetOptions(options []string) {
	f.options = options
}

func (f *FieldAbstract) Options() []string {
	return f.options
}