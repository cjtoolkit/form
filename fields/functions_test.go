package fields

import (
	"github.com/cjtoolkit/form"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var panicChannel = make(chan interface{})

func panicTrap(fn func()) {
	defer func() {
		panicChannel <- recover()
	}()
	fn()
}

func TestFunctions(t *testing.T) {
	Convey("UseDefaultKeyIfCustomKeyIsEmpty", t, func() {

		Convey("Use defaultKey because customKey is empty", func() {
			So(UseDefaultKeyIfCustomKeyIsEmpty("default", ""), ShouldEqual, "default")
		})

		Convey("Use customKey because customKey is not empty", func() {
			So(UseDefaultKeyIfCustomKeyIsEmpty("default", "custom"), ShouldEqual, "custom")
		})

	})

	Convey("ExecFuncIfErrIsNotNil", t, func() {

		Convey("Err is nil, therefore won't execute function", func() {
			execute := false

			So(ExecFuncIfErrIsNotNil(nil, func() {
				execute = true
			}), ShouldEqual, false)

			So(execute, ShouldEqual, false)
		})

		Convey("Err is not nil, therefore execute function", func() {
			execute := false

			So(ExecFuncIfErrIsNotNil(form.ErrorUnknown("Hi!"), func() {
				execute = true
			}), ShouldEqual, true)

			So(execute, ShouldEqual, true)
		})

	})
}
