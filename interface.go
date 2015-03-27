package form

// Interface
type Interface interface {
	CJForm(*Fields)
}

// StructPtr Hijacker Interface
type Hijacker interface {
	CJStructPtr() interface{}
}
