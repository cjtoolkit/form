package form

// Interface
type StructPtrForm interface {
	CJForm(*Fields)
}

// StructPtr Hijacker Interface
type Hijacker interface {
	CJStructPtr() interface{}
}
