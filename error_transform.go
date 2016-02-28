package form

/*
Implement:
	error in "builtin"
*/
type ErrorTransform string

func (eT ErrorTransform) Error() string {
	return string(eT)
}
