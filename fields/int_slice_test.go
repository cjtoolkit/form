package fields

import (
	"github.com/cjtoolkit/form"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestIntSlice(t *testing.T) {
	form.FormFieldInterfaceCheck(IntSlice{})

	Convey("PreCheck", t, func() {

		Convey("Panic because Name is empty", func() {
			go panicTrap(func() {
				(IntSlice{}).PreCheck()
			})

			So(<-panicChannel, ShouldEqual,
				form.ErrorPreCheck("IntSlice Field: Name cannot be empty string"))
		})

		Convey("Panic because Label is empty", func() {
			go panicTrap(func() {
				(IntSlice{
					Name: "test",
				}).PreCheck()
			})

			So(<-panicChannel, ShouldEqual,
				form.ErrorPreCheck("IntSlice Field: test: Label cannot be empty string"))
		})

		Convey("Panic because Norm is nil", func() {
			go panicTrap(func() {
				(IntSlice{
					Name:  "test",
					Label: "test",
				}).PreCheck()
			})

			So(<-panicChannel, ShouldEqual,
				form.ErrorPreCheck("IntSlice Field: test: Norm cannot be nil value"))
		})

		Convey("Panic because Model is nil", func() {
			var norm []string

			go panicTrap(func() {
				(IntSlice{
					Name:  "test",
					Label: "test",
					Norm:  &norm,
				}).PreCheck()
			})

			So(<-panicChannel, ShouldEqual,
				form.ErrorPreCheck("IntSlice Field: test: Model cannot be nil value"))
		})

		Convey("Panic because Err is nil", func() {
			var norm []string
			var model []int64

			go panicTrap(func() {
				(IntSlice{
					Name:  "test",
					Label: "test",
					Norm:  &norm,
					Model: &model,
				}).PreCheck()
			})

			So(<-panicChannel, ShouldEqual,
				form.ErrorPreCheck("IntSlice Field: test: Err cannot be nil value"))
		})

		Convey("Everything checks out", func() {
			var norm []string
			var model []int64
			var err error

			go panicTrap(func() {
				(IntSlice{
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

	Convey("ReverseTransform", t, func() {

		Convey("Panic because a few are not integer", func() {
			norm := []string{"3", "a", "1"}
			var model []int64

			go panicTrap(func() {
				(IntSlice{
					Norm:  &norm,
					Model: &model,
				}).ReverseTransform()
			})

			So(<-panicChannel, ShouldResemble, &form.ErrorReverseTransform{
				Key: form.LANG_NOT_INT,
				Value: IntSlice{
					Norm:  &norm,
					Model: &model,
				},
			})

			So(model, ShouldResemble, []int64{3})
		})

		Convey("Should not panic because all are integer", func() {
			norm := []string{"3", "2", "1"}
			var model []int64

			go panicTrap(func() {
				(IntSlice{
					Norm:  &norm,
					Model: &model,
				}).ReverseTransform()
			})

			So(<-panicChannel, ShouldBeNil)

			So(model, ShouldResemble, []int64{1, 2, 3})
		})

	})

	Convey("ValidateModel", t, func() {

		Convey("validateRequired", func() {

			Convey("Does nothing because it's not required", func() {
				var model []int64

				go panicTrap(func() {
					(IntSlice{
						Model: &model,
					}).validateRequired()
				})

				So(<-panicChannel, ShouldBeNil)
			})

			Convey("Panic because model is either empty or nil while being required", func() {
				var model []int64

				go panicTrap(func() {
					(IntSlice{
						Model:    &model,
						Required: true,
					}).validateRequired()
				})

				So(<-panicChannel, ShouldResemble, &form.ErrorValidateModel{
					Key: form.LANG_FIELD_REQUIRED,
					Value: IntSlice{
						Model:    &model,
						Required: true,
					},
				})

				model = []int64{}

				go panicTrap(func() {
					(IntSlice{
						Model:    &model,
						Required: true,
					}).validateRequired()
				})

				So(<-panicChannel, ShouldResemble, &form.ErrorValidateModel{
					Key: form.LANG_FIELD_REQUIRED,
					Value: IntSlice{
						Model:    &model,
						Required: true,
					},
				})
			})

			Convey("Should not panic because model is neither empty or nil while being required", func() {
				var model []int64 = []int64{5}

				go panicTrap(func() {
					(IntSlice{
						Model:    &model,
						Required: true,
					}).validateRequired()
				})

				So(<-panicChannel, ShouldBeNil)
			})

		})

	})
}
