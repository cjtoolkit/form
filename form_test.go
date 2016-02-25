package form

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestForm(t *testing.T) {
	// start let
	form := &Form{}
	// end let

	Convey("checkValues", t, func() {

		Convey("Panic if 'values' is 'nil'", func() {
			defer func() {
				So(recover(), ShouldEqual,
					"SetForm or SetMultipartForm has not been called or nil value has been passed to either.")
			}()

			form.checkValues()
		})

		Convey("'values' is not nil, therefore should not panic", func() {
			defer func() {
				So(recover(), ShouldBeNil)
			}()

			form.values = &values{}

			form.checkValues()

			form.values = &valuesFile{}

			form.checkValues()

			form.values = nil
		})

	})

	Convey("handleError", t, func() {

		Convey("check all three case", func() {
			var err interface{}
			var output error

			checkError := func(err interface{}) {
				defer form.handleError(&output)
				panic(err)
			}

			err = &ErrorTransform{}
			checkError(err)
			So(output, ShouldEqual, err)

			err = ErrorUnknown("Hi")
			checkError(err)
			So(output, ShouldEqual, err)

			err = "Hello, World!"
			checkError(err)
			So(output, ShouldEqual, ErrorUnknown("Hello, World!"))
		})

	})

	Convey("checkErrorInLoop", t, func() {

		Convey("'success' should stay 'true' if 'err' is 'nil'", func() {
			success := true

			form.checkErrorInLoop(nil, &success)

			So(success, ShouldEqual, true)
		})

		Convey("'success' should change to 'false' if 'err' is not 'nil'", func() {
			success := true

			form.checkErrorInLoop(ErrorUnknown("Hi"), &success)

			So(success, ShouldEqual, false)
		})

	})

	Convey("checkForm", t, func() {

		Convey("Panic if 'form' is 'nil'", func() {
			defer func() {
				So(recover(), ShouldEqual, "'form' cannot be nil")
			}()

			form.checkForm(nil)
		})

		Convey("'form' is not 'nil', therefore does not panic", func() {
			defer func() {
				So(recover(), ShouldBeNil)
			}()

			form.checkForm(&fakeForm{})
		})

	})

	Convey("checkField", t, func() {

		Convey("Panic if 'field' is 'nil'", func() {
			defer func() {
				So(recover(), ShouldEqual, "'field' cannot be nil")
			}()

			form.checkField(nil)
		})

		Convey("'field' is not 'nil', therefore does not panic", func() {
			defer func() {
				So(recover(), ShouldBeNil)
			}()

			form.checkField(&fakeFormField{})
		})

	})

	Convey("checkErrPtr", t, func() {

		Convey("Panic if 'errPtr' is 'nil'", func() {
			defer func() {
				So(recover(), ShouldEqual, "'errPtr' cannot be nil")
			}()

			form.checkErrPtr(nil)
		})

		Convey("'errPtr' is not 'nil', therefore does not panic", func() {
			defer func() {
				So(recover(), ShouldBeNil)
			}()

			var err error

			form.checkErrPtr(&err)
		})

	})
}
