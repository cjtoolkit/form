// Foundation Second Layer (http://foundation.zurb.com/)
package foundation

import (
	"fmt"
	"github.com/cjtoolkit/form"
	"io"
	"strings"
)

type foundation struct{}

func (f foundation) addClass(src string, classes ...string) (str string) {
	if src == "" {
		str = strings.Join(classes, " ")
		return
	}
	srcs := strings.Split(src, " ")
	srcs = append(srcs, classes...)
	str = strings.Join(srcs, " ")
	return
}

func (f foundation) Render(w io.Writer, r form.RenderData) {
	before := ""
	after := ""

	r.Fns.Call("html", map[string]interface{}{
		"before": &before,
		"after":  &after,
	})

	fmt.Fprint(w, before)

	fmt.Fprintf(w, `<div id="form-group-%d">`, r.Count)

	beforeInput := ""
	afterInput := ""
	startOfGroup := ""
	endOfGroup := ""

	r.Fns.Call("foundation", map[string]interface{}{
		"beforeInput":  &beforeInput,
		"afterInput":   &afterInput,
		"startOfGroup": &startOfGroup,
		"endOfGroup":   &endOfGroup,
	})

	fmt.Fprint(w, startOfGroup)

	labelContent := ""
	labelFor := ""
	var labelAttr map[string]string
	renderedLabelAttr := ""

	if r.Type == form.InputRadio {
		goto formField
	}

	r.Fns.Call("label", map[string]interface{}{
		"content": &labelContent,
		"for":     &labelFor,
		"attr":    &labelAttr,
	})

	if labelFor == "" {
		goto formField
	}

	if labelAttr != nil {
		delete(labelAttr, "for")
		if r.Error != nil {
			labelAttr["class"] = f.addClass(labelAttr["class"], "error")
		}
		renderedLabelAttr = form.RenderAttr(labelAttr)
	}

	fmt.Fprintf(w, `<label for="%s" %s>%s</label>`, es(labelFor), renderedLabelAttr, es(labelContent))

formField:

	fmt.Fprint(w, beforeInput)

	// render first layer stacks
	for _, field := range r.FirstLayerStacks {
		switch field := field.(type) {
		case *form.FirstLayerInput:
			addErrorAndRender := func() {
				if r.Error != nil {
					field.Attr["class"] = f.addClass(field.Attr["class"], "error")
				}
				field.Render(w)
			}
			switch field.Attr["type"] {
			case "radio":
				label := strings.TrimSpace(field.Label)
				if label != "" {
					field.Label = ""
					if r.Error == nil {
						fmt.Fprint(w, `<label>`)
					} else {
						fmt.Fprint(w, `<label class="error">`)
					}
					addErrorAndRender()
					fmt.Fprintf(w, ` %s</label>`, es(label))
				} else {
					addErrorAndRender()
				}

			default:
				addErrorAndRender()
			}
		case *form.FirstLayerTextarea:
			if r.Error != nil {
				field.Attr["class"] = f.addClass(field.Attr["class"], "error")
			}
			field.Render(w)
		case *form.FirstLayerSelect:
			if r.Error != nil {
				field.Attr["class"] = f.addClass(field.Attr["class"], "error")
			}
			field.Render(w)
		}
	}

	fmt.Fprint(w, afterInput)

	if r.Error != nil {
		fmt.Fprintf(w, `<small class="error">%s</small>`, es(r.Error.Error()))
	} else {
		fmt.Fprint(w, `<small class="error" style="display: none;"></small>`)
	}

	fmt.Fprint(w, endOfGroup)

	fmt.Fprint(w, `</div>`) // Close form group

	fmt.Fprint(w, after)
}

/*
Foundation Second Layer (http://foundation.zurb.com/)
*/
func SecondLayer() form.RenderSecondLayer {
	return foundation{}
}
