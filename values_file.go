package form

import (
	"mime/multipart"
)

/*
Implement:
	ValuesInterface
	ValuesFileInterface
*/
type valuesFile struct {
	values      *multipart.Form
	valuesCount map[string]int
	fileCount   map[string]int
}

func newValuesFile(form *multipart.Form) *valuesFile {
	return &valuesFile{
		values:      form,
		valuesCount: map[string]int{},
		fileCount:   map[string]int{},
	}
}

func (v *valuesFile) valueIncrement(name string) {
	v.valuesCount[name]++
}

func (v *valuesFile) GetOne(name string) string {
	if nil == v.values.Value[name] || v.valuesCount[name] >= len(v.values.Value[name]) {
		return ""
	}
	defer v.valueIncrement(name)
	return v.values.Value[name][v.valuesCount[name]]
}

func (v *valuesFile) valueMarkAll(name string) {
	v.valuesCount[name] = len(v.values.Value[name])
}

func (v *valuesFile) GetAll(name string) []string {
	if nil == v.values.Value[name] || v.valuesCount[name] >= len(v.values.Value[name]) {
		return nil
	}
	defer v.valueMarkAll(name)
	return v.values.Value[name][v.valuesCount[name]:]
}

func (v *valuesFile) fileIncrement(name string) {
	v.fileCount[name]++
}

func (v *valuesFile) GetOneFile(name string) *multipart.FileHeader {
	if nil == v.values.File[name] || v.fileCount[name] >= len(v.values.File[name]) {
		return nil
	}
	defer v.fileIncrement(name)
	return v.values.File[name][v.fileCount[name]]
}

func (v *valuesFile) fileMarkAll(name string) {
	v.fileCount[name] = len(v.values.File[name])
}

func (v *valuesFile) GetAllFile(name string) []*multipart.FileHeader {
	if nil == v.values.File[name] || v.fileCount[name] >= len(v.values.File[name]) {
		return nil
	}
	defer v.fileMarkAll(name)
	return v.values.File[name][v.fileCount[name]:]
}
