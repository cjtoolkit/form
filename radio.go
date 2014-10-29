package form

// For use with InputRadio (string)
type Radio struct {
	Value    string
	Label    string
	Selected bool
	Attr     map[string]string
}

// For use with InputRadio (int64)
type RadioInt struct {
	Value    int64
	Label    string
	Selected bool
	Attr     map[string]string
}

// For use with InputRadio (float64)
type RadioFloat struct {
	Value    float64
	Label    string
	Selected bool
	Attr     map[string]string
}
