package fields

type Int64Sort []int64

func (p Int64Sort) Len() int           { return len(p) }
func (p Int64Sort) Less(i, j int) bool { return p[i] < p[j] }
func (p Int64Sort) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
