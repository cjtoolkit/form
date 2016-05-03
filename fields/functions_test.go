package fields

import (
	"fmt"
	"github.com/cjtoolkit/form"
	. "github.com/smartystreets/goconvey/convey"
	"regexp"
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

	Convey("getMessageFromError", t, func() {

		Convey("Err is nil, return empty string", func() {
			So(getMessageFromError(nil), ShouldBeEmpty)
		})

		Convey("Err is not nil, return message", func() {
			So(getMessageFromError(fmt.Errorf("Hi")), ShouldEqual, "Hi")
		})

	})

	Convey("getPatternFromRegExp", t, func() {

		Convey("Re is nil, return empty string", func() {
			So(getPatternFromRegExp(nil), ShouldBeEmpty)
		})

		Convey("Re is not nil, return pattern", func() {
			So(getPatternFromRegExp(regexp.MustCompile(`[a-z]`)), ShouldEqual, `[a-z]`)
		})

	})

	Convey("addSuffix", t, func() {

		name := "name"
		suffix := "suffix"

		So(addSuffix(name, nil), ShouldEqual, "name")
		So(addSuffix(name, &suffix), ShouldEqual, "name-suffix")

	})
}
