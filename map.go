package form

type StringMap map[string]bool

func StringSliceToMap(a []string) StringMap {
	m := StringMap{}
	for _, v := range a {
		m[v] = true
	}
	return m
}

func (s StringMap) Index(v string) bool {
	return s[v]
}

type IntMap map[int64]bool

func IntSliceToMap(a []int64) IntMap {
	m := IntMap{}
	for _, v := range a {
		m[v] = true
	}
	return m
}

func (i IntMap) Index(v int64) bool {
	return i[v]
}