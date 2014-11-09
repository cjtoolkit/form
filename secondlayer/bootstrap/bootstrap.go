// Bootstrap Second Layer (http://getbootstrap.com/)
package bootstrap

import (
	"fmt"
	"github.com/cjtoolkit/form"
	"io"
)

/*
Bootstrap Second Layer (http://getbootstrap.com/)
	"html":
		"before": *string
		"after": *string
	"label":
		"content": *string
		"for": *string
		"attr": *map[string]string
	"bootstrap":
		"beforeInput": *string
		"afterInput": *string
		"startOfGroup": *string
		"endOfGroup": *string
		"helpBlock": *string
		"disabled": *bool (Only works with 'InputRadio' and 'InputCheckbox')
		"feedback" *bool
*/
func SecondLayer() form.RenderSecondLayer {
	return bootstrap{}
}

type bootstrap struct{}

func (b bootstrap) addClass(src string, classes ...string) (str string) {
	if src == "" {
		str = strings.Join(classes, " ")
		return
	}
	srcs := strings.Split(src, " ")
	srcs = append(srcs, classes...)
	str = strings.Join(srcs, " ")
	return
}

func (b bootstrap) Render(w io.Writer, r form.RenderData) {
	if r.Type == form.InputHidden {
		r.FirstLayerStacks.Render(w)
		return
	}

	beforeInput := ""
	afterInput := ""
	startOfGroup := ""
	endOfGroup := ""
	helpBlock := ""
	disabled := false
	feedback := false

	r.Fns.Call("bootstrap", map[string]interface{}{
		"beforeInput":  &beforeInput,
		"afterInput":   &afterInput,
		"startOfGroup": &startOfGroup,
		"endOfGroup":   &endOfGroup,
		"helpBlock":    &helpBlock,
		"disabled":     &disabled,
		"feedback":     &feedback,
	})

	before := ""
	after := ""

	r.Fns.Call("html", map[string]interface{}{
		"before": &before,
		"after":  &after,
	})

	fmt.Fprint(w, before)

	eclass := ""

	if r.Error != nil {
		eclass += "has-error "
		if feedback {
			eclass += "has-feedback"
		}
	} else if r.Warning != "" {
		eclass += "has-warning "
		if feedback {
			eclass += "has-feedback"
		}
	} else if r.Check {
		eclass += "has-success "
		if feedback {
			eclass += "has-feedback"
		}
	}

	strDisabled := ""
	if disabled {
		strDisabled = " disabled"
	}

	switch r.Type {
	case form.InputRadio:
		fmt.Fprintf(w, `<div id="form-group-%d" class="%s"><div class="radio%s">`,
			r.Count, eclass, strDisabled)
	case form.InputCheckbox:
		fmt.Fprintf(w, `<div id="form-group-%d" class="%s"><div class="checkbox%s">`,
			r.Count, eclass, strDisabled)
	default:
		fmt.Fprintf(w, `<div id="form-group-%d" class="form-group %s">`, r.Count, eclass)
	}

	fmt.Fprint(w, startOfGroup)

	labelContent := ""
	labelFor := ""
	var labelAttr map[string]string
	parsedLabelAttr := ""

	if r.Type == form.InputRadio {
		goto formField
	}

	r.Fns.Call("label", map[string]interface{}{
		"content": &labelContent,
		"for":     &labelFor,
		"attr":    &labelAttr,
	})

	if r.Type == form.InputCheckbox || r.Type == form.InputRadio {
		fmt.Fprint(w, `<label `)
		if labelAttr != nil {
			delete(labelAttr, "for")
			fmt.Fprint(w, form.RenderAttr(labelAttr))
		}
		fmt.Fprint(w, `>`)
		goto formField
	} else if labelFor == "" {
		goto formField
	}

	if labelAttr != nil {
		delete(labelAttr, "for")
		parsedLabelAttr = form.RenderAttr(labelAttr)
	}

	fmt.Fprintf(w, `<label for="%s" %s>%s</label>`, es(labelFor), parsedLabelAttr, es(labelContent))

formField:

	fmt.Fprint(w, beforeInput)

	for _, field := range r.FirstLayerStacks {
		switch field := field.(type) {
		case *form.FirstLayerInput:
			switch field.Attr["type"] {
			case "radio", "checkbox", "file", "range", "hidden", "color":
				// Do nothing
			default:
				field.Attr["class"] = b.addClass(field.Attr["class"], "form-control")
			}
		case *form.FirstLayerSelect:
			field.Attr["class"] = b.addClass(field.Attr["class"], "form-control")
		case *form.FirstLayerTextarea:
			field.Attr["class"] = b.addClass(field.Attr["class"], "form-control")
		}
		field.Render(w)
	}

	fmt.Fprint(w, afterInput)

	if r.Type == form.InputCheckbox || r.Type == form.InputRadio {
		fmt.Fprint(w, es(labelContent), `</label>`)
	}

	if r.Error != nil {
		fmt.Fprintf(w, `<span class="help-block">%s</span>`, es(r.Error.Error()))
	} else if r.Warning != "" {
		fmt.Fprintf(w, `<span class="help-block">%s</span>`, es(r.Warning))
	} else {
		fmt.Fprintf(w, `<span class="help-block">%s</span>`, es(helpBlock))
	}

	fmt.Fprint(w, endOfGroup)

	switch r.Type {
	case form.InputRadio, form.InputCheckbox:
		fmt.Fprint(w, `</div></div>`) // Close form group
	default:
		fmt.Fprint(w, `</div>`) // Close form group
	}

	fmt.Fprint(w, after)
}
