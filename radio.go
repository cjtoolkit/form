package form

type Radio struct {
	Value    string
	Label    string
	Selected bool
	Attr     map[string]string
}

type RadioInt struct {
	Value    int64
	Label    string
	Selected bool
	Attr     map[string]string
}

type RadioFloat struct {
	Value    float64
	Label    string
	Selected bool
	Attr     map[string]string
}
