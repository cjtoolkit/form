package data

import (
	"io"

	"github.com/cjtoolkit/form"
)

// DataSlice Interface
type DataSlice interface {
	form.RenderSecondLayer

	// Pull and Reset Data
	PullAndReset() []Data
}

// Create new Data Slice
//
// Note: Do not use as Default Second Renderer, not thread safe.
func NewDataSlice() DataSlice {
	return &dataSlice{}
}

type dataSlice []Data

func (d *dataSlice) Render(w io.Writer, r form.RenderData) {
	data := &Data{}

	data.Name = r.Name
	data.Count = r.Count

	if r.Error != nil {
		data.Error = r.Error.Error()
	}

	data.Warning = r.Warning
	data.Checked = r.Check
	data.FirstLayerStacks = r.FirstLayerStacks

	labelContent := ""
	labelFor := ""
	var labelAttr map[string]string

	r.Fns.Call("label", map[string]interface{}{
		"content": &labelContent,
		"for":     &labelFor,
		"attr":    &labelAttr,
	})

	data.LabelFor = labelFor
	data.LabelContent = labelContent

	if labelAttr != nil {
		data.LabelAttr = form.RenderAttr(labelAttr)
	}

	*d = append(*d, *data)
}

func (d *dataSlice) PullAndReset() []Data {
	sliceData := []Data(*d)
	*d = dataSlice{}
	return sliceData
}
