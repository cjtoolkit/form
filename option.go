package form

// For use with Select (string, []string)
type Option struct {
	Content  string
	Value    string
	Label    string
	Selected bool
	Attr     map[string]string
}

// For use with Select (int64, []int64)
type OptionInt struct {
	Content  string
	Value    int64
	Label    string
	Selected bool
	Attr     map[string]string
}

// For use with Select (float64, []float64)
type OptionFloat struct {
	Content  string
	Value    float64
	Label    string
	Selected bool
	Attr     map[string]string
}
