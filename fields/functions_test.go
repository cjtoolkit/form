package fields

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestFunctions(t *testing.T) {
	Convey("UseDefaultKeyIfCustomKeyIsEmpty", t, func() {

		Convey("Use defaultKey because customKey is empty", func() {
			So(UseDefaultKeyIfCustomKeyIsEmpty("default", ""), ShouldEqual, "default")
		})

		Convey("Use customKey because customKey is not empty", func() {
			So(UseDefaultKeyIfCustomKeyIsEmpty("default", "custom"), ShouldEqual, "custom")
		})

	})
}
