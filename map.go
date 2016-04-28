package form

type StringMap map[string]bool

// Convert String Slice to Map
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

// Convert Int Slice to Map
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
