package form

import (
	"net/url"
)

/*
Implement:
	ValuesInterface
*/
type values struct {
	values url.Values
	counts map[string]int
}

func newValues(_values url.Values) *values {
	return &values{
		values: _values,
		counts: map[string]int{},
	}
}

func (v *values) increment(name string) {
	v.counts[name]++
}

func (v *values) GetOne(name string) string {
	if nil == v.values[name] || v.counts[name] >= len(v.values[name]) {
		return ""
	}
	defer v.increment(name)
	return v.values[name][v.counts[name]]
}

func (v *values) markAll(name string) {
	v.counts[name] = len(v.values[name])
}

func (v *values) GetAll(name string) []string {
	if nil == v.values[name] || v.counts[name] >= len(v.values[name]) {
		return nil
	}
	defer v.markAll(name)
	return v.values[name][v.counts[name]:]
}
