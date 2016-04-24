package fields

var panicChannel = make(chan interface{})

func panicTrap(fn func()) {
	defer func() {
		panicChannel <- recover()
	}()
	fn()
}
