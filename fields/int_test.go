package fields

import (
	"github.com/cjtoolkit/form"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestInt(t *testing.T) {
	form.FormFieldInterfaceCheck(Int{})

	Convey("PreCheck", t, func() {

		Convey("Should panic because Name is empty string", func() {
			go panicTrap(func() { (Int{}).PreCheck() })

			So(<-panicChannel, ShouldEqual, form.ErrorPreCheck("Int Field: Name cannot be empty string"))
		})

		Convey("Should panic because Label is empty string", func() {
			go panicTrap(func() {
				(Int{
					Name: "hello",
				}).PreCheck()
			})

			So(<-panicChannel, ShouldEqual, form.ErrorPreCheck("Int Field: hello: Label cannot be empty string"))
		})

		Convey("Should panic because Norm is nil value", func() {
			go panicTrap(func() {
				(Int{
					Name:  "hello",
					Label: "Hello",
				}).PreCheck()
			})

			So(<-panicChannel, ShouldEqual, form.ErrorPreCheck("Int Field: hello: Norm cannot be nil value"))
		})

		Convey("Should panic because Model is nil value", func() {
			var norm string

			go panicTrap(func() {
				(Int{
					Name:  "hello",
					Label: "Hello",
					Norm:  &norm,
				}).PreCheck()
			})

			So(<-panicChannel, ShouldEqual, form.ErrorPreCheck("Int Field: hello: Model cannot be nil value"))
		})

		Convey("Should panic because Err is nil value", func() {
			var norm string
			var model int64

			go panicTrap(func() {
				(Int{
					Name:  "hello",
					Label: "Hello",
					Norm:  &norm,
					Model: &model,
				}).PreCheck()
			})

			So(<-panicChannel, ShouldEqual, form.ErrorPreCheck("Int Field: hello: Err cannot be nil value"))
		})

		Convey("Every mandatory field is in order, so therefore should not panic", func() {
			var norm string
			var model int64
			var err error

			go panicTrap(func() {
				(Int{
					Name:  "hello",
					Label: "Hello",
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

			Convey("Should not panic because field is not required", func() {
				go panicTrap(func() { (Int{}).validateRequired() })

				So(<-panicChannel, ShouldBeNil)
			})

			Convey("Should panic because required field is an empty string", func() {
				model := int64(0)

				go panicTrap(func() {
					(Int{
						Model:    &model,
						Required: true,
					}).validateRequired()
				})

				So(<-panicChannel, ShouldResemble, &form.ErrorValidateModel{
					Key: form.LANG_FIELD_REQUIRED,
					Value: Int{
						Model:    &model,
						Required: true,
					},
				})
			})

			Convey("Should not panic because required field is not an empty string", func() {
				model := int64(5)

				go panicTrap(func() {
					(Int{
						Model:    &model,
						Required: true,
					}).validateRequired()
				})

				So(<-panicChannel, ShouldBeNil)
			})

		})

		Convey("validateMin", func() {

			Convey("Don't validate 0 if MinZero is false", func() {
				model := int64(-5)

				go panicTrap(func() {
					(Int{
						Model: &model,
					}).validateMin()
				})

				So(<-panicChannel, ShouldBeNil)
			})

			Convey("Validate 0 if MinZero is true, should panic because it's less than 0", func() {
				model := int64(-5)

				go panicTrap(func() {
					(Int{
						Model:   &model,
						MinZero: true,
					}).validateMin()
				})

				So(<-panicChannel, ShouldResemble, &form.ErrorValidateModel{
					Key: form.LANG_NUMBER_MIN,
					Value: Int{
						Model:   &model,
						MinZero: true,
					},
				})
			})

			Convey("Validate 0 if MinZero is true, shoud not panic because it's more than 0", func() {
				model := int64(5)

				go panicTrap(func() {
					(Int{
						Model:   &model,
						MinZero: true,
					}).validateMin()
				})

				So(<-panicChannel, ShouldBeNil)
			})

			Convey("Validate 5, should panic because it's less than 5", func() {
				model := int64(2)

				go panicTrap(func() {
					(Int{
						Model: &model,
						Min:   5,
					}).validateMin()
				})

				So(<-panicChannel, ShouldResemble, &form.ErrorValidateModel{
					Key: form.LANG_NUMBER_MIN,
					Value: Int{
						Model: &model,
						Min:   5,
					},
				})
			})

			Convey("Validate 5, should not panic because it's more than 5", func() {
				model := int64(6)

				go panicTrap(func() {
					(Int{
						Model: &model,
						Min:   5,
					}).validateMin()
				})

				So(<-panicChannel, ShouldBeNil)
			})

		})

		Convey("validateMax", func() {

			Convey("Don't validate 0 if MaxZero is false", func() {
				model := int64(-5)

				go panicTrap(func() {
					(Int{
						Model: &model,
					}).validateMax()
				})

				So(<-panicChannel, ShouldBeNil)
			})

			Convey("Validate 0 if MaxZero is true, should panic because it's more than 0", func() {
				model := int64(5)

				go panicTrap(func() {
					(Int{
						Model:   &model,
						MaxZero: true,
					}).validateMax()
				})

				So(<-panicChannel, ShouldResemble, &form.ErrorValidateModel{
					Key: form.LANG_NUMBER_MAX,
					Value: Int{
						Model:   &model,
						MaxZero: true,
					},
				})
			})

			Convey("Validate 0 if MaxZero is true, shoud not panic because it's less than 0", func() {
				model := int64(-5)

				go panicTrap(func() {
					(Int{
						Model:   &model,
						MaxZero: true,
					}).validateMax()
				})

				So(<-panicChannel, ShouldBeNil)
			})

			Convey("Validate 5, should panic because it's more than 5", func() {
				model := int64(6)

				go panicTrap(func() {
					(Int{
						Model: &model,
						Max:   5,
					}).validateMax()
				})

				So(<-panicChannel, ShouldResemble, &form.ErrorValidateModel{
					Key: form.LANG_NUMBER_MAX,
					Value: Int{
						Model: &model,
						Max:   5,
					},
				})
			})

			Convey("Validate 5, should not panic because it's less than 5", func() {
				model := int64(2)

				go panicTrap(func() {
					(Int{
						Model: &model,
						Max:   5,
					}).validateMax()
				})

				So(<-panicChannel, ShouldBeNil)
			})

		})

		Convey("validateStep", func() {

			Convey("Do nothing because step is set to Zero", func() {
				go panicTrap(func() { (Int{}).validateStep() })

				So(<-panicChannel, ShouldBeNil)
			})

			Convey("Should because panic because model is not in step", func() {
				model := int64(3)

				go panicTrap(func() {
					(Int{
						Model: &model,
						Step:  2,
					}).validateStep()
				})

				So(<-panicChannel, ShouldResemble, &form.ErrorValidateModel{
					Key: form.LANG_NUMBER_STEP,
					Value: Int{
						Model: &model,
						Step:  2,
					},
				})
			})

			Convey("Should not panic because model is in step", func() {
				model := int64(4)

				go panicTrap(func() {
					(Int{
						Model: &model,
						Step:  2,
					}).validateStep()
				})

				So(<-panicChannel, ShouldBeNil)
			})

		})

		Convey("validateInList", func() {

			Convey("Should not panic, because InList is 'nil'", func() {
				go panicTrap(func() { (Int{}).validateInList() })

				So(<-panicChannel, ShouldBeNil)
			})

			Convey("Should not panic, because Model is in the List", func() {
				model := int64(42)
				list := []int64{12, 42, 60}

				go panicTrap(func() {
					(Int{
						Model:  &model,
						InList: list,
					}).validateInList()
				})

				So(<-panicChannel, ShouldBeNil)
			})

			Convey("Should panic, because Model is not in the List", func() {
				model := int64(50)
				list := []int64{12, 42, 60}

				go panicTrap(func() {
					(Int{
						Model:  &model,
						InList: list,
					}).validateInList()
				})

				So(<-panicChannel, ShouldResemble, &form.ErrorValidateModel{
					Key: form.LANG_IN_LIST,
					Value: Int{
						Model:  &model,
						InList: list,
					},
				})
			})

		})

	})
}
