package fields

import (
	"github.com/cjtoolkit/form"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestText(t *testing.T) {
	Convey("PreCheck", t, func() {
		Convey("Panic because Name is empty string", func() {
			defer func() {
				So(recover(), ShouldEqual, form.ErrorPreCheck("Text Field: Name cannot be empty string"))
			}()

			(Text{}).PreCheck()
		})

		Convey("Panic because Label is empty string", func() {
			defer func() {
				So(recover(), ShouldEqual, form.ErrorPreCheck("Text Field: hello: Label cannot be empty string"))
			}()

			(Text{
				Name: "hello",
			}).PreCheck()
		})

		Convey("Panic because Norm is nil value", func() {
			defer func() {
				So(recover(), ShouldEqual, form.ErrorPreCheck("Text Field: hello: Norm cannot be nil value"))
			}()

			(Text{
				Name:  "hello",
				Label: "Hello",
			}).PreCheck()
		})

		Convey("Panic because Model is nil value", func() {
			defer func() {
				So(recover(), ShouldEqual, form.ErrorPreCheck("Text Field: hello: Model cannot be nil value"))
			}()

			var norm string

			(Text{
				Name:  "hello",
				Label: "Hello",
				Norm:  &norm,
			}).PreCheck()
		})

		Convey("Panic because Err is nil value", func() {
			defer func() {
				So(recover(), ShouldEqual, form.ErrorPreCheck("Text Field: hello: Err cannot be nil value"))
			}()

			var norm, model string

			(Text{
				Name:  "hello",
				Label: "Hello",
				Norm:  &norm,
				Model: &model,
			}).PreCheck()
		})

		Convey("Every mandatory field is in order, so therefore should not panic", func() {
			defer func() {
				So(recover(), ShouldBeNil)
			}()

			var norm, model string
			var err error

			(Text{
				Name:  "hello",
				Label: "Hello",
				Norm:  &norm,
				Model: &model,
				Err:   &err,
			}).PreCheck()
		})
	})

	Convey("ValidateModel", t, func() {

		Convey("validateRequired", func() {

			Convey("Should not panic because field is not required", func() {
				defer func() {
					So(recover(), ShouldBeNil)
				}()

				(Text{}).validateRequired()
			})

			Convey("Panic because required field is an empty string", func() {
				defer func() {
					So(recover(), ShouldResemble, &form.ErrorValidateModel{
						Key: form.LANG_FIELD_REQUIRED,
						Value: map[string]interface{}{
							"Label": "test",
						},
					})
				}()

				model := ""

				(Text{
					Label:    "test",
					Model:    &model,
					Required: true,
				}).validateRequired()
			})

			Convey("Should not panic because required field is not an empty string", func() {
				defer func() {
					So(recover(), ShouldBeNil)
				}()

				model := "hello"

				(Text{
					Label:    "test",
					Model:    &model,
					Required: true,
				}).validateRequired()
			})

		})

		Convey("validateMinChar", func() {

			Convey("Should not panic because MinChar has not been populated", func() {
				defer func() {
					So(recover(), ShouldBeNil)
				}()

				(Text{}).validateMinChar()
			})

			Convey("Panic because model is less than MinChar", func() {
				defer func() {
					So(recover(), ShouldResemble, &form.ErrorValidateModel{
						Key: form.LANG_MIN_CHAR,
						Value: map[string]interface{}{
							"Label":   "test",
							"MinChar": 4,
						},
					})
				}()

				model := "he"

				(Text{
					Label:   "test",
					MinChar: 4,
					Model:   &model,
				}).validateMinChar()
			})

			Convey("Should not panic because model is more than MinChar", func() {
				defer func() {
					So(recover(), ShouldBeNil)
				}()

				model := "hello"

				(Text{
					Label:   "test",
					MinChar: 4,
					Model:   &model,
				}).validateMinChar()
			})

		})

		Convey("validateMaxChar", func() {

			Convey("Should not panic because MaxChar has not been populated", func() {
				defer func() {
					So(recover(), ShouldBeNil)
				}()

				(Text{}).validateMaxChar()
			})

			Convey("Panic because model is greater than MaxChar", func() {
				defer func() {
					So(recover(), ShouldResemble, &form.ErrorValidateModel{
						Key: form.LANG_MAX_CHAR,
						Value: map[string]interface{}{
							"Label":   "test",
							"MaxChar": 4,
						},
					})
				}()

				model := "hello"

				(Text{
					Label:   "test",
					MaxChar: 4,
					Model:   &model,
				}).validateMaxChar()
			})

			Convey("Should not panic because model is less tahn MinChar", func() {
				defer func() {
					So(recover(), ShouldBeNil)
				}()

				model := "he"

				(Text{
					Label:   "test",
					MaxChar: 4,
					Model:   &model,
				}).validateMaxChar()
			})

		})

	})
}
