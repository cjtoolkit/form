package form

/*
Implement:
	error in "builtin"
*/
type ErrorUnknown string

func (eU ErrorUnknown) Error() string {
	return string(eU)
}
