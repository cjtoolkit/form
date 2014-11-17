/*
Data Harvest.


Usage
	f := formStr{}
	dataSlice := data.NewDataSlice()
	check := form.New(dataSlice, "en-GB")
	check.RenderStr(&f)
	d := dataSlice.PullAndReset()
*/
package data

import (
	"bytes"
	"html/template"

	"github.com/cjtoolkit/form"
)

// Data
type Data struct {
	Name             string
	Count            int
	Error            string
	Warning          string
	Checked          bool
	FirstLayerStacks form.FirstLayerStack
	LabelFor         string
	LabelContent     string
	LabelAttr        string // Pre Rendered
}

// Render First Layer Stacks as 'string'
func (d Data) RenderString() string {
	b := &bytes.Buffer{}
	defer b.Reset()
	d.FirstLayerStacks.Render(b)
	return b.String()
}

// Render First Layer Stacks as 'template.HTML'
func (d Data) RenderHtml() template.HTML {
	return template.HTML(d.RenderString())
}
