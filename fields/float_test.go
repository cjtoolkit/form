package fields

import (
	"github.com/cjtoolkit/form"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestFloat(t *testing.T) {
	Convey("PreCheck", t, func() {

		Convey("Should panic because Name is empty string", func() {
			go panicTrap(func() {
				(Float{}).PreCheck()
			})

			So(<-panicChannel, ShouldEqual, form.ErrorPreCheck("Float Field: Name cannot be empty string"))
		})

		Convey("Should panic because Label is empty string", func() {
			go panicTrap(func() {
				(Float{
					Name: "hello",
				}).PreCheck()
			})

			So(<-panicChannel, ShouldEqual, form.ErrorPreCheck("Float Field: hello: Label cannot be empty string"))
		})

		Convey("Should panic because Norm is nil value", func() {
			go panicTrap(func() {
				(Float{
					Name:  "hello",
					Label: "Hello",
				}).PreCheck()
			})

			So(<-panicChannel, ShouldEqual, form.ErrorPreCheck("Float Field: hello: Norm cannot be nil value"))
		})

		Convey("Should panic because Model is nil value", func() {
			var norm string

			go panicTrap(func() {
				(Float{
					Name:  "hello",
					Label: "Hello",
					Norm:  &norm,
				}).PreCheck()
			})

			So(<-panicChannel, ShouldEqual, form.ErrorPreCheck("Float Field: hello: Model cannot be nil value"))
		})

		Convey("Should panic because Err is nil value", func() {
			var norm string
			var model float64

			go panicTrap(func() {
				(Float{
					Name:  "hello",
					Label: "Hello",
					Norm:  &norm,
					Model: &model,
				}).PreCheck()
			})

			So(<-panicChannel, ShouldEqual, form.ErrorPreCheck("Float Field: hello: Err cannot be nil value"))
		})

		Convey("Every mandatory field is in order, so therefore should not panic", func() {
			var norm string
			var model float64
			var err error

			go panicTrap(func() {
				(Float{
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
				go panicTrap(func() {
					(Float{}).validateRequired()
				})

				So(<-panicChannel, ShouldBeNil)
			})

			Convey("Should panic because required field is an empty string", func() {
				model := float64(0)

				go panicTrap(func() {
					(Float{
						Label:    "required",
						Model:    &model,
						Required: true,
					}).validateRequired()
				})

				So(<-panicChannel, ShouldResemble, &form.ErrorValidateModel{
					Key: form.LANG_FIELD_REQUIRED,
					Value: map[string]interface{}{
						"Label": "required",
					},
				})
			})

			Convey("Should not panic because required field is not an empty string", func() {
				model := float64(5)

				go panicTrap(func() {
					(Float{
						Label:    "required",
						Model:    &model,
						Required: true,
					}).validateRequired()
				})

				So(<-panicChannel, ShouldBeNil)
			})

		})

		Convey("validateMin", func() {

			Convey("Don't validate 0 if MinZero is false", func() {
				label := "min"
				model := float64(-5)

				go panicTrap(func() {
					(Float{
						Label: label,
						Model: &model,
					}).validateMin()
				})

				So(<-panicChannel, ShouldBeNil)
			})

			Convey("Validate 0 if MinZero is true, should panic because it's less than 0", func() {
				label := "min"
				model := float64(-0.1)

				go panicTrap(func() {
					(Float{
						Label:   label,
						Model:   &model,
						MinZero: true,
					}).validateMin()
				})

				So(<-panicChannel, ShouldResemble, &form.ErrorValidateModel{
					Key: form.LANG_NUMBER_MIN,
					Value: map[string]interface{}{
						"Label": "min",
						"Min":   float64(0),
					},
				})
			})

			Convey("Validate 0 if MinZero is true, shoud not panic because it's more than 0", func() {
				label := "min"
				model := float64(0.1)

				go panicTrap(func() {
					(Float{
						Label:   label,
						Model:   &model,
						MinZero: true,
					}).validateMin()
				})

				So(<-panicChannel, ShouldBeNil)
			})

			Convey("Validate 5, should panic because it's less than 5", func() {
				label := "min"
				model := float64(4.9)

				go panicTrap(func() {
					(Float{
						Label: label,
						Model: &model,
						Min:   5,
					}).validateMin()
				})

				So(<-panicChannel, ShouldResemble, &form.ErrorValidateModel{
					Key: form.LANG_NUMBER_MIN,
					Value: map[string]interface{}{
						"Label": "min",
						"Min":   float64(5),
					},
				})
			})

			Convey("Validate 5, should not panic because it's more than 5", func() {
				label := "min"
				model := float64(5.1)

				go panicTrap(func() {
					(Float{
						Label: label,
						Model: &model,
						Min:   5,
					}).validateMin()
				})

				So(<-panicChannel, ShouldBeNil)
			})

		})

		Convey("validateMax", func() {

			Convey("Don't validate 0 if MaxZero is false", func() {
				label := "max"
				model := float64(-5)

				go panicTrap(func() {
					(Float{
						Label: label,
						Model: &model,
					}).validateMax()
				})

				So(<-panicChannel, ShouldBeNil)
			})

			Convey("Validate 0 if MaxZero is true, should panic because it's more than 0", func() {
				label := "max"
				model := float64(0.1)

				go panicTrap(func() {
					(Float{
						Label:   label,
						Model:   &model,
						MaxZero: true,
					}).validateMax()
				})

				So(<-panicChannel, ShouldResemble, &form.ErrorValidateModel{
					Key: form.LANG_NUMBER_MAX,
					Value: map[string]interface{}{
						"Label": "max",
						"Max":   float64(0),
					},
				})
			})

			Convey("Validate 0 if MaxZero is true, shoud not panic because it's less than 0", func() {
				label := "max"
				model := float64(-0.1)

				go panicTrap(func() {
					(Float{
						Label:   label,
						Model:   &model,
						MaxZero: true,
					}).validateMax()
				})

				So(<-panicChannel, ShouldBeNil)
			})

			Convey("Validate 5, should panic because it's more than 5", func() {
				label := "max"
				model := float64(5.1)

				go panicTrap(func() {
					(Float{
						Label: label,
						Model: &model,
						Max:   5,
					}).validateMax()
				})

				So(<-panicChannel, ShouldResemble, &form.ErrorValidateModel{
					Key: form.LANG_NUMBER_MAX,
					Value: map[string]interface{}{
						"Label": "max",
						"Max":   float64(5),
					},
				})
			})

			Convey("Validate 5, should not panic because it's less than 5", func() {
				label := "min"
				model := float64(4.9)

				go panicTrap(func() {
					(Float{
						Label: label,
						Model: &model,
						Max:   5,
					}).validateMax()
				})

				So(<-panicChannel, ShouldBeNil)
			})

		})

		Convey("validateStep", func() {

			Convey("Do nothing because step is set to Zero", func() {

				model := float64(3)

				go panicTrap(func() {
					(Float{Model: &model}).validateStep()
				})

				So(<-panicChannel, ShouldBeNil)
			})

			Convey("Should because panic because model is not in step", func() {

				label := "step"
				model := float64(3)

				go panicTrap(func() {
					(Float{
						Label: label,
						Model: &model,
						Step:  2,
					}).validateStep()
				})

				So(<-panicChannel, ShouldResemble, &form.ErrorValidateModel{
					Key: form.LANG_NUMBER_STEP,
					Value: map[string]interface{}{
						"Label": "step",
						"Step":  float64(2),
					},
				})
			})

			Convey("Should not panic because model is in step", func() {

				label := "step"
				model := float64(4)

				go panicTrap(func() {
					(Float{
						Label: label,
						Model: &model,
						Step:  2,
					}).validateStep()
				})

				So(<-panicChannel, ShouldBeNil)
			})

		})

		Convey("validateInList", func() {

			Convey("Should not panic, because InList is 'nil'", func() {

				go panicTrap(func() {
					(Float{}).validateInList()
				})

				So(<-panicChannel, ShouldBeNil)
			})

			Convey("Should not panic, because Model is in the List", func() {

				label := "List"
				model := float64(1.5)
				list := []float64{1.4, 1.5, 1.6}

				go panicTrap(func() {
					(Float{
						Label:  label,
						Model:  &model,
						InList: list,
					}).validateInList()
				})

				So(<-panicChannel, ShouldBeNil)
			})

			Convey("Should panic, because Model is not in the List", func() {

				label := "List"
				model := float64(1.3)
				list := []float64{1.4, 1.5, 1.6}

				go panicTrap(func() {
					(Float{
						Label:  label,
						Model:  &model,
						InList: list,
					}).validateInList()
				})

				So(<-panicChannel, ShouldResemble, &form.ErrorValidateModel{
					Key: form.LANG_IN_LIST,
					Value: map[string]interface{}{
						"Label": "List",
						"List":  []float64{1.4, 1.5, 1.6},
					},
				})
			})

		})

	})
}
