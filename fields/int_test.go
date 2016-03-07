package fields

import (
	"github.com/cjtoolkit/form"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestInt(t *testing.T) {
	Convey("PreCheck", t, func() {
		Convey("Should panic because Name is empty string", func() {
			defer func() {
				So(recover(), ShouldEqual, form.ErrorPreCheck("Int Field: Name cannot be empty string"))
			}()

			(Int{}).PreCheck()
		})

		Convey("Should panic because Label is empty string", func() {
			defer func() {
				So(recover(), ShouldEqual, form.ErrorPreCheck("Int Field: hello: Label cannot be empty string"))
			}()

			(Int{
				Name: "hello",
			}).PreCheck()
		})

		Convey("Should panic because Norm is nil value", func() {
			defer func() {
				So(recover(), ShouldEqual, form.ErrorPreCheck("Int Field: hello: Norm cannot be nil value"))
			}()

			(Int{
				Name:  "hello",
				Label: "Hello",
			}).PreCheck()
		})

		Convey("Should panic because Model is nil value", func() {
			defer func() {
				So(recover(), ShouldEqual, form.ErrorPreCheck("Int Field: hello: Model cannot be nil value"))
			}()

			var norm string

			(Int{
				Name:  "hello",
				Label: "Hello",
				Norm:  &norm,
			}).PreCheck()
		})

		Convey("Should panic because Err is nil value", func() {
			defer func() {
				So(recover(), ShouldEqual, form.ErrorPreCheck("Int Field: hello: Err cannot be nil value"))
			}()

			var norm string
			var model int64

			(Int{
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

			var norm string
			var model int64
			var err error

			(Int{
				Name:  "hello",
				Label: "Hello",
				Norm:  &norm,
				Model: &model,
				Err:   &err,
			}).PreCheck()
		})
	})

	Convey("ValidateModel", t, func() {

		Convey("validateMin", func() {

			Convey("Don't validate 0 if MinZero is false", func() {
				defer func() {
					So(recover(), ShouldBeNil)
				}()

				label := "min"
				model := int64(-5)

				(Int{
					Label: label,
					Model: &model,
				}).validateMin()
			})

			Convey("Validate 0 if MinZero is true, should panic because it's less than 0", func() {
				defer func() {
					So(recover(), ShouldResemble, &form.ErrorValidateModel{
						Key: form.LANG_NUMBER_MIN,
						Value: map[string]interface{}{
							"Label": "min",
							"Min":   int64(0),
						},
					})
				}()

				label := "min"
				model := int64(-5)

				(Int{
					Label:   label,
					Model:   &model,
					MinZero: true,
				}).validateMin()
			})

			Convey("Validate 0 if MinZero is true, shoud not panic because it's more than 0", func() {
				defer func() {
					So(recover(), ShouldBeNil)
				}()

				label := "min"
				model := int64(5)

				(Int{
					Label:   label,
					Model:   &model,
					MinZero: true,
				}).validateMin()
			})

			Convey("Validate 5, should panic because it's less than 5", func() {
				defer func() {
					So(recover(), ShouldResemble, &form.ErrorValidateModel{
						Key: form.LANG_NUMBER_MIN,
						Value: map[string]interface{}{
							"Label": "min",
							"Min":   int64(5),
						},
					})
				}()

				label := "min"
				model := int64(2)

				(Int{
					Label: label,
					Model: &model,
					Min:   5,
				}).validateMin()
			})

			Convey("Validate 5, should not panic because it's more than 5", func() {
				defer func() {
					So(recover(), ShouldBeNil)
				}()

				label := "min"
				model := int64(6)

				(Int{
					Label: label,
					Model: &model,
					Min:   5,
				}).validateMin()
			})

		})

	})
}
