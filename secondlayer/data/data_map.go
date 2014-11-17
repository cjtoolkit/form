package data

import (
	"io"

	"github.com/cjtoolkit/form"
)

// DataMap Interface
type DataMap interface {
	form.RenderSecondLayer

	// Pull and Reset Data
	PullAndReset() map[string][]Data
}

// Create New Data Map.
//
// Note: Do not use as Default Second Renderer, not thread safe.
func NewDataMap() DataMap {
	return &dataMap{}
}

type dataMap map[string][]Data

func (d *dataMap) Render(w io.Writer, r form.RenderData) {
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

	m := *d
	m[r.Name] = append(m[r.Name], *data)
}

func (d *dataMap) PullAndReset() map[string][]Data {
	mapData := map[string][]Data(*d)
	*d = dataMap{}
	return mapData
}
