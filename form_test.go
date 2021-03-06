package form

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestForm(t *testing.T) {
	// start let
	form := &Form{}
	// end let

	Convey("checkLanguage", t, func() {

		Convey("Panic if 'language' is nil", func() {
			defer func() {
				So(recover(), ShouldEqual, "Form language cannot be set to nil")
			}()

			form.checkLanguage()
		})

		Convey("'language' is not nil, therefore should not panic", func() {
			defer func() {
				So(recover(), ShouldBeNil)
			}()

			form.language = Langauge{}

			form.checkLanguage()

			form.language = nil
		})

	})

	Convey("checkValues", t, func() {

		Convey("Panic if 'values' is 'nil'", func() {
			defer func() {
				So(recover(), ShouldEqual,
					"SetForm or SetMultipartForm cannot be nil value.")
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

		form.language = Langauge{}

		Convey("check all five case", func() {
			var err interface{}
			var errPtr error

			checkError := func(err interface{}) {
				defer form.handleError(&errPtr)
				panic(err)
			}

			func() {
				err = &ErrorValidateModel{}

				defer func() {
					So(recover(), ShouldBeNil)
					So(errPtr, ShouldEqual, err)
				}()

				checkError(err)
			}()

			func() {
				err = ErrorUnknown("Hi")

				defer func() {
					So(recover(), ShouldBeNil)
					So(errPtr, ShouldEqual, err)
				}()

				errPtr = nil
				checkError(err)
			}()

			func() {
				defer func() {
					So(recover(), ShouldBeNil)
					So(errPtr, ShouldEqual, ErrorUnknown("Hello, World!"))
				}()

				errPtr = nil
				err = "Hello, World!"
				checkError(err)
			}()

			func() {
				errPtr = nil
				defer func() {
					So(recover(), ShouldBeNil)
					So(errPtr, ShouldBeNil)
				}()

				err = nil
				checkError(err)

			}()

			func() {
				errPtr = nil
				defer func() {
					So(recover(), ShouldEqual, ErrorTransform("hello"))
					So(errPtr, ShouldBeNil)
				}()

				err = ErrorTransform("hello")
				checkError(err)
			}()
		})

		form.language = nil

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
