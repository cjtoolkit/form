package fields

import (
	"github.com/cjtoolkit/form"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestStringSlice(t *testing.T) {
	form.FormFieldInterfaceCheck(StringSlice{})

	Convey("PreCheck", t, func() {

		Convey("Panic because Name is empty", func() {
			go panicTrap(func() {
				(StringSlice{}).PreCheck()
			})

			So(<-panicChannel, ShouldEqual,
				form.ErrorPreCheck("StringSlice Field: Name cannot be empty string"))
		})

		Convey("Panic because Label is empty", func() {
			go panicTrap(func() {
				(StringSlice{
					Name: "test",
				}).PreCheck()
			})

			So(<-panicChannel, ShouldEqual,
				form.ErrorPreCheck("StringSlice Field: test: Label cannot be empty string"))
		})

		Convey("Panic because Norm is nil", func() {
			go panicTrap(func() {
				(StringSlice{
					Name:  "test",
					Label: "test",
				}).PreCheck()
			})

			So(<-panicChannel, ShouldEqual,
				form.ErrorPreCheck("StringSlice Field: test: Norm cannot be nil value"))
		})

		Convey("Panic because Model is nil", func() {
			var norm []string

			go panicTrap(func() {
				(StringSlice{
					Name:  "test",
					Label: "test",
					Norm:  &norm,
				}).PreCheck()
			})

			So(<-panicChannel, ShouldEqual,
				form.ErrorPreCheck("StringSlice Field: test: Model cannot be nil value"))
		})

		Convey("Panic because Err is nil", func() {
			var norm []string
			var model []string

			go panicTrap(func() {
				(StringSlice{
					Name:  "test",
					Label: "test",
					Norm:  &norm,
					Model: &model,
				}).PreCheck()
			})

			So(<-panicChannel, ShouldEqual,
				form.ErrorPreCheck("StringSlice Field: test: Err cannot be nil value"))
		})

		Convey("Everything checks out", func() {
			var norm []string
			var model []string
			var err error

			go panicTrap(func() {
				(StringSlice{
					Name:  "test",
					Label: "test",
					Norm:  &norm,
					Model: &model,
					Err:   &err,
				}).PreCheck()
			})

			So(<-panicChannel, ShouldBeNil)
		})

	})

	Convey("ValidateModel", t, func() {

		Convey("validateRequired", func() {

			Convey("Does nothing because it's not required", func() {
				var model []string

				go panicTrap(func() {
					(StringSlice{
						Model: &model,
					}).validateRequired()
				})

				So(<-panicChannel, ShouldBeNil)
			})

			Convey("Panic because model is either empty or nil while being required", func() {
				var model []string

				go panicTrap(func() {
					(StringSlice{
						Model:    &model,
						Required: true,
					}).validateRequired()
				})

				So(<-panicChannel, ShouldResemble, &form.ErrorValidateModel{
					Key: form.LANG_FIELD_REQUIRED,
					Value: StringSlice{
						Model:    &model,
						Required: true,
					},
				})

				model = []string{}

				go panicTrap(func() {
					(StringSlice{
						Model:    &model,
						Required: true,
					}).validateRequired()
				})

				So(<-panicChannel, ShouldResemble, &form.ErrorValidateModel{
					Key: form.LANG_FIELD_REQUIRED,
					Value: StringSlice{
						Model:    &model,
						Required: true,
					},
				})
			})

			Convey("Should not panic because model is neither empty or nil while being required", func() {
				var model []string = []string{"Hey"}

				go panicTrap(func() {
					(StringSlice{
						Model:    &model,
						Required: true,
					}).validateRequired()
				})

				So(<-panicChannel, ShouldBeNil)
			})

		})

	})
}
