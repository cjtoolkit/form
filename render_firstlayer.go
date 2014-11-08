package form

import (
	"fmt"
	"io"
)

// FirstLayer Interface
type FirstLayer interface {
	Render(w io.Writer)
}

// First Layer Stack
type FirstLayerStack []FirstLayer

func (f *FirstLayerStack) append(e FirstLayer) {
	*f = append(*f, e)
}

// Render
func (f FirstLayerStack) Render(w io.Writer) {
	for _, item := range f {
		item.Render(w)
	}
}

// First Layer Input
type FirstLayerInput struct {
	Attr  map[string]string
	Label string
}

// Render
func (f *FirstLayerInput) Render(w io.Writer) {
	if f.Label == "" {
		fmt.Fprintf(w, `<input %s/>`, RenderAttr(f.Attr))
	} else {
		fmt.Fprintf(w, `<label><input %s/>%s</label>`, RenderAttr(f.Attr), es(f.Label))
	}
}

// First Layer Textarea
type FirstLayerTextarea struct {
	Attr    map[string]string
	Content string
}

// Render
func (f *FirstLayerTextarea) Render(w io.Writer) {
	fmt.Fprintf(w, `<textarea %s>%s</textare>`, RenderAttr(f.Attr), es(f.Content))
}

// First Layer Select
type FirstLayerSelect struct {
	Attr    map[string]string
	Options []Option
}

// Render
func (f *FirstLayerSelect) Render(w io.Writer) {
	fmt.Fprintf(w, `<select %s>`, RenderAttr(f.Attr))
	for _, option := range f.Options {
		delete(option.Attr, "value")
		fmt.Fprintf(w, `<option value="%s" %s>%s</option>`,
			es(option.Value), RenderAttr(option.Attr), es(option.Content))
	}
	fmt.Fprint(w, `</select>`)
}
