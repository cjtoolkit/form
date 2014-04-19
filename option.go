package form

type Option struct {
	Content  string
	Value    string
	Label    string
	Selected bool
	Attr     map[string]string
}

type OptionInt struct {
	Content  string
	Value    int64
	Label    string
	Selected bool
	Attr     map[string]string
}

type OptionFloat struct {
	Content  string
	Value    float64
	Label    string
	Selected bool
	Attr     map[string]string
}
