package form

func FormDisablePreCheck() func(*Form) {
	return func(f *Form) {
		f.disablePreCheck = true
	}
}
