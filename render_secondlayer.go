package form

import (
	"fmt"
	"io"
)

// For use with RenderSecondLayer
type RenderData struct {
	Name             string
	Count            int
	Type             TypeCode
	Error            error
	Warning          string
	Fns              FieldFuncs
	Check            bool
	FirstLayerStacks FirstLayerStack
}

// Second Layer interface
type RenderSecondLayer interface {
	Render(w io.Writer, r RenderData)
}

// An function adapter for 'RenderSecondLayer'
type RenderSecondLayerFunc func(w io.Writer, r RenderData)

// A method that just calls the function.
func (render RenderSecondLayerFunc) Render(w io.Writer, r RenderData) {
	render(w, r)
}

// Default Second Layer
var DefaultRenderSecondLayer RenderSecondLayer = RenderSecondLayerFunc(defaultRenderSecondLayer)

func defaultRenderSecondLayer(w io.Writer, r RenderData) {
	if r.Type == InputHidden {
		r.FirstLayerStacks.Render(w)
		return
	}

	before := ""
	after := ""

	r.Fns.Call("html", map[string]interface{}{
		"before": &before,
		"after":  &after,
	})

	fmt.Fprint(w, before)

	fmt.Fprintf(w, `<div id="form-group-%d" class="form-group">`, r.Count)

	labelContent := ""
	labelFor := ""
	var labelAttr map[string]string
	parsedLabelAttr := ""

	if r.Type == InputRadio {
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
		parsedLabelAttr = RenderAttr(labelAttr)
	}

	fmt.Fprintf(w, `<label for="%s" %s>%s</label>`, es(labelFor), parsedLabelAttr, es(labelContent))

formField:

	r.FirstLayerStacks.Render(w)

	if r.Error != nil {
		fmt.Fprintf(w, `<div class="error">%s</div>`, es(r.Error.Error()))
	} else {
		fmt.Fprint(w, `<div class="error" style="display: none;"></div>`)
	}

	if r.Warning != "" {
		fmt.Fprintf(w, `<div class="warning">%s</div>`, es(r.Warning))
	} else {
		fmt.Fprint(w, `<div class="warning" style="display: none;"></div>`)
	}

	fmt.Fprint(w, `</div>`) // Close form group

	fmt.Fprint(w, after)
}
