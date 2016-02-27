package form

/*
Implement
	error in "builtin"
*/
type ErrorPreCheck string

func (ePC ErrorPreCheck) Error() string {
	return string(ePC)
}
