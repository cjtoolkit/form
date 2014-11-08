package form

import (
	"fmt"
)

func (r renderValue) strSelect(value string) {
	_select := &FirstLayerSelect{Options: []Option{}}
	r.fls.append(_select)

	var attr map[string]string

	r.fieldsFns.Call("attr", map[string]interface{}{
		"attr": &attr,
	})

	if attr != nil {
		delete(attr, "name")
		delete(attr, "multiple")
		_select.Attr = attr
	} else {
		_select.Attr = map[string]string{}
	}

	var options []Option

	r.fieldsFns.Call("option", map[string]interface{}{
		"option": &options,
	})

	if options == nil {
		panic(fmt.Errorf(r.form.T("ErrSelectNotWellFormed")))
		return
	}

	_select.Attr["name"] = r.preferedName

	for _, option := range options {
		_option := &Option{}

		if option.Value != "" {
			_option.Value = option.Value
		}

		_option.Content = option.Content

		if option.Attr != nil {
			delete(option.Attr, "value")
			delete(option.Attr, "label")
			delete(option.Attr, "selected")
			_option.Attr = option.Attr
		} else {
			_option.Attr = map[string]string{}
		}

		if option.Label != "" {
			_option.Attr["label"] = option.Label
		}

		if value == "" {
			if option.Selected {
				_option.Attr["selected"] = " "
			}
		} else {
			if value == option.Value {
				_option.Attr["selected"] = " "
			}
		}

		_select.Options = append(_select.Options, *_option)
	}
}

func (r renderValue) strsSelect(values []string) {
	_select := &FirstLayerSelect{Options: []Option{}}
	r.fls.append(_select)

	var attr map[string]string

	r.fieldsFns.Call("attr", map[string]interface{}{
		"attr": &attr,
	})

	if attr != nil {
		delete(attr, "name")
		delete(attr, "multiple")
		_select.Attr = attr
	} else {
		_select.Attr = map[string]string{}
	}

	var options []Option

	r.fieldsFns.Call("option", map[string]interface{}{
		"option": &options,
	})

	if options == nil {
		panic(fmt.Errorf(r.form.T("ErrSelectNotWellFormed")))
		return
	}

	_select.Attr["name"] = r.preferedName
	_select.Attr["multiple"] = " "

	for _, option := range options {
		_option := &Option{}

		if option.Value != "" {
			_option.Value = option.Value
		}

		_option.Content = option.Content

		if option.Attr != nil {
			delete(option.Attr, "value")
			delete(option.Attr, "label")
			delete(option.Attr, "selected")
			_option.Attr = option.Attr
		} else {
			_option.Attr = map[string]string{}
		}

		if option.Label != "" {
			_option.Attr["label"] = option.Label
		}

		if len(values) == 0 {
			if option.Selected {
				_option.Attr["selected"] = " "
			}
		} else {
			for _, value := range values {
				if value == option.Value {
					_option.Attr["selected"] = " "
					break
				}
			}
		}

		_select.Options = append(_select.Options, *_option)
	}
}

func (r renderValue) wnumSelect(value int64) {
	_select := &FirstLayerSelect{Options: []Option{}}
	r.fls.append(_select)

	var attr map[string]string

	r.fieldsFns.Call("attr", map[string]interface{}{
		"attr": &attr,
	})

	if attr != nil {
		delete(attr, "name")
		delete(attr, "multiple")
		_select.Attr = attr
	} else {
		_select.Attr = map[string]string{}
	}

	var options []OptionInt

	r.fieldsFns.Call("option", map[string]interface{}{
		"option": &options,
	})

	if options == nil {
		panic(fmt.Errorf(r.form.T("ErrSelectNotWellFormed")))
		return
	}

	_select.Attr["name"] = r.preferedName

	for _, option := range options {
		_option := &Option{}

		if option.Value != 0 {
			_option.Value = fmt.Sprintf("%d", option.Value)
		}

		_option.Content = option.Content

		if option.Attr != nil {
			delete(option.Attr, "value")
			delete(option.Attr, "label")
			delete(option.Attr, "selected")
			_option.Attr = option.Attr
		} else {
			_option.Attr = map[string]string{}
		}

		if option.Label != "" {
			_option.Attr["label"] = option.Label
		}

		if value == 0 {
			if option.Selected {
				_option.Attr["selected"] = " "
			}
		} else {
			if value == option.Value {
				_option.Attr["selected"] = " "
			}
		}

		_select.Options = append(_select.Options, *_option)
	}
}

func (r renderValue) wnumsSelect(values []int64) {
	_select := &FirstLayerSelect{Options: []Option{}}
	r.fls.append(_select)

	var attr map[string]string

	r.fieldsFns.Call("attr", map[string]interface{}{
		"attr": &attr,
	})

	if attr != nil {
		delete(attr, "name")
		delete(attr, "multiple")
		_select.Attr = attr
	} else {
		_select.Attr = map[string]string{}
	}

	var options []OptionInt

	r.fieldsFns.Call("option", map[string]interface{}{
		"option": &options,
	})

	if options == nil {
		panic(fmt.Errorf(r.form.T("ErrSelectNotWellFormed")))
		return
	}

	_select.Attr["name"] = r.preferedName
	_select.Attr["multiple"] = " "

	for _, option := range options {
		_option := &Option{}

		if option.Value != 0 {
			_option.Value = fmt.Sprintf("%d", option.Value)
		}

		_option.Content = option.Content

		if option.Attr != nil {
			delete(option.Attr, "value")
			delete(option.Attr, "label")
			delete(option.Attr, "selected")
			_option.Attr = option.Attr
		} else {
			_option.Attr = map[string]string{}
		}

		if option.Label != "" {
			_option.Attr["label"] = option.Label
		}

		if len(values) == 0 {
			if option.Selected {
				_option.Attr["selected"] = " "
			}
		} else {
			for _, value := range values {
				if value == option.Value {
					_option.Attr["selected"] = " "
					break
				}
			}
		}
	}
}

func (r renderValue) fnumSelect(value float64) {
	_select := &FirstLayerSelect{Options: []Option{}}
	r.fls.append(_select)

	var attr map[string]string

	r.fieldsFns.Call("attr", map[string]interface{}{
		"attr": &attr,
	})

	if attr != nil {
		delete(attr, "name")
		delete(attr, "multiple")
		_select.Attr = attr
	} else {
		_select.Attr = map[string]string{}
	}

	var options []OptionFloat

	r.fieldsFns.Call("option", map[string]interface{}{
		"option": &options,
	})

	if options == nil {
		panic(fmt.Errorf(r.form.T("ErrSelectNotWellFormed")))
		return
	}

	_select.Attr["name"] = r.preferedName

	for _, option := range options {
		_option := &Option{}

		if option.Value != 0 {
			_option.Value = fmt.Sprintf("%d", option.Value)
		}

		_option.Content = option.Content

		if option.Attr != nil {
			delete(option.Attr, "value")
			delete(option.Attr, "label")
			delete(option.Attr, "selected")
			_option.Attr = option.Attr
		} else {
			_option.Attr = map[string]string{}
		}

		if option.Label != "" {
			_option.Attr["label"] = option.Label
		}

		if value == 0 {
			if option.Selected {
				_option.Attr["selected"] = " "
			}
		} else {
			if value == option.Value {
				_option.Attr["selected"] = " "
			}
		}

		_select.Options = append(_select.Options, *_option)
	}
}

func (r renderValue) fnumsSelect(values []float64) {
	_select := &FirstLayerSelect{Options: []Option{}}
	r.fls.append(_select)

	var attr map[string]string

	r.fieldsFns.Call("attr", map[string]interface{}{
		"attr": &attr,
	})

	if attr != nil {
		delete(attr, "name")
		delete(attr, "multiple")
		_select.Attr = attr
	} else {
		_select.Attr = map[string]string{}
	}

	var options []OptionFloat

	r.fieldsFns.Call("option", map[string]interface{}{
		"option": &options,
	})

	if options == nil {
		panic(fmt.Errorf(r.form.T("ErrSelectNotWellFormed")))
		return
	}

	_select.Attr["name"] = r.preferedName
	_select.Attr["multiple"] = " "

	for _, option := range options {
		_option := &Option{}

		if option.Value != 0 {
			_option.Value = fmt.Sprintf("%d", option.Value)
		}

		_option.Content = option.Content

		if option.Attr != nil {
			delete(option.Attr, "value")
			delete(option.Attr, "label")
			delete(option.Attr, "selected")
			_option.Attr = option.Attr
		} else {
			_option.Attr = map[string]string{}
		}

		if option.Label != "" {
			_option.Attr["label"] = option.Label
		}

		if len(values) == 0 {
			if option.Selected {
				_option.Attr["selected"] = " "
			}
		} else {
			for _, value := range values {
				if value == option.Value {
					_option.Attr["selected"] = " "
					break
				}
			}
		}
	}
}
