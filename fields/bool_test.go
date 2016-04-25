package fields

import (
	"github.com/cjtoolkit/form"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestBool(t *testing.T) {
	Convey("PreCheck", t, func() {

		Convey("Panic because name is empty", func() {
			go panicTrap(func() {
				(Bool{}).PreCheck()
			})

			So(<-panicChannel, ShouldEqual, form.ErrorPreCheck("Bool Field: Name cannot be empty string"))
		})

		Convey("Panic because label is empty", func() {
			go panicTrap(func() {
				(Bool{
					Name: "bool",
				}).PreCheck()
			})

			So(<-panicChannel, ShouldEqual,
				form.ErrorPreCheck("Bool Field: bool: Label cannot be empty string"))
		})

		Convey("Panic because norm is nil", func() {
			go panicTrap(func() {
				(Bool{
					Name:  "bool",
					Label: "bool",
				}).PreCheck()
			})

			So(<-panicChannel, ShouldEqual,
				form.ErrorPreCheck("Bool Field: bool: Norm cannot be nil value"))
		})

		Convey("Panic because model is nil", func() {
			var norm string

			go panicTrap(func() {
				(Bool{
					Name:  "bool",
					Label: "bool",
					Norm:  &norm,
				}).PreCheck()
			})

			So(<-panicChannel, ShouldEqual,
				form.ErrorPreCheck("Bool Field: bool: Model cannot be nil value"))
		})

		Convey("Panic because err is nil", func() {
			var norm string
			var model bool

			go panicTrap(func() {
				(Bool{
					Name:  "bool",
					Label: "bool",
					Norm:  &norm,
					Model: &model,
				}).PreCheck()
			})

			So(<-panicChannel, ShouldEqual,
				form.ErrorPreCheck("Bool Field: bool: Err cannot be nil value"))
		})

		Convey("Panic because value is empty", func() {
			var norm string
			var model bool
			var err error

			go panicTrap(func() {
				(Bool{
					Name:  "bool",
					Label: "bool",
					Norm:  &norm,
					Model: &model,
					Err:   &err,
				}).PreCheck()
			})

			So(<-panicChannel, ShouldEqual,
				form.ErrorPreCheck("Bool Field: bool: Value cannot be empty string"))
		})

		Convey("Everything checks out", func() {
			var norm string
			var model bool
			var err error

			go panicTrap(func() {
				(Bool{
					Name:  "bool",
					Label: "bool",
					Norm:  &norm,
					Model: &model,
					Err:   &err,
					Value: "hi",
				}).PreCheck()
			})

			So(<-panicChannel, ShouldBeNil)
		})

	})

	Convey("Transform", t, func() {

		Convey("Model is false, therefore does not update norm", func() {
			norm := ""
			model := false

			(Bool{
				Norm:  &norm,
				Model: &model,
				Value: "hi",
			}).Transform()

			So(norm, ShouldBeEmpty)
		})

		Convey("Model is true, there does update norm with value", func() {
			norm := ""
			model := true

			(Bool{
				Norm:  &norm,
				Model: &model,
				Value: "hi",
			}).Transform()

			So(norm, ShouldEqual, "hi")
		})

	})

	Convey("ReverseTransform", t, func() {

		Convey("Norm and Value do not match, therefore Model should be false", func() {
			norm := ""
			var model bool

			(Bool{
				Norm:  &norm,
				Model: &model,
				Value: "hi",
			}).ReverseTransform()

			So(model, ShouldBeFalse)
		})

		Convey("Norm and Value does match, therefore Model should be true", func() {
			norm := "hi"
			var model bool

			(Bool{
				Norm:  &norm,
				Model: &model,
				Value: "hi",
			}).ReverseTransform()

			So(model, ShouldBeTrue)
		})

	})

	Convey("validateRequired", t, func() {

		Convey("Do nothing because Required has not been specified", func() {
			model := false

			go panicTrap(func() {
				(Bool{
					Model: &model,
				}).validateRequired()
			})

			So(<-panicChannel, ShouldBeNil)
		})

		Convey("Panic because Model is false while is required", func() {
			model := false

			go panicTrap(func() {
				(Bool{
					Model:    &model,
					Required: true,
				}).validateRequired()
			})

			So(<-panicChannel, ShouldResemble, &form.ErrorValidateModel{
				Key: form.LANG_FIELD_REQUIRED,
				Value: map[string]interface{}{
					"Label": "",
				},
			})
		})

		Convey("Should not panic because Model is true while is required", func() {
			model := true

			go panicTrap(func() {
				(Bool{
					Model:    &model,
					Required: true,
				}).validateRequired()
			})

			So(<-panicChannel, ShouldBeNil)
		})

	})
}
