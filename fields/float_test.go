package fields

import (
	"github.com/cjtoolkit/form"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestFloat(t *testing.T) {
	Convey("PreCheck", t, func() {

		Convey("Should panic because Name is empty string", func() {
			defer func() {
				So(recover(), ShouldEqual, form.ErrorPreCheck("Float Field: Name cannot be empty string"))
			}()

			(Float{}).PreCheck()
		})

		Convey("Should panic because Label is empty string", func() {
			defer func() {
				So(recover(), ShouldEqual, form.ErrorPreCheck("Float Field: hello: Label cannot be empty string"))
			}()

			(Float{
				Name: "hello",
			}).PreCheck()
		})

		Convey("Should panic because Norm is nil value", func() {
			defer func() {
				So(recover(), ShouldEqual, form.ErrorPreCheck("Float Field: hello: Norm cannot be nil value"))
			}()

			(Float{
				Name:  "hello",
				Label: "Hello",
			}).PreCheck()
		})

		Convey("Should panic because Model is nil value", func() {
			defer func() {
				So(recover(), ShouldEqual, form.ErrorPreCheck("Float Field: hello: Model cannot be nil value"))
			}()

			var norm string

			(Float{
				Name:  "hello",
				Label: "Hello",
				Norm:  &norm,
			}).PreCheck()
		})

		Convey("Should panic because Err is nil value", func() {
			defer func() {
				So(recover(), ShouldEqual, form.ErrorPreCheck("Float Field: hello: Err cannot be nil value"))
			}()

			var norm string
			var model float64

			(Float{
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
			var model float64
			var err error

			(Float{
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

				(Float{}).validateRequired()
			})

			Convey("Should panic because required field is an empty string", func() {
				defer func() {
					So(recover(), ShouldResemble, &form.ErrorValidateModel{
						Key: form.LANG_FIELD_REQUIRED,
						Value: map[string]interface{}{
							"Label": "required",
						},
					})
				}()

				model := float64(0)

				(Float{
					Label:    "required",
					Model:    &model,
					Required: true,
				}).validateRequired()
			})

			Convey("Should not panic because required field is not an empty string", func() {
				defer func() {
					So(recover(), ShouldBeNil)
				}()

				model := float64(5)

				(Float{
					Label:    "required",
					Model:    &model,
					Required: true,
				}).validateRequired()
			})

		})

		Convey("validateMin", func() {

			Convey("Don't validate 0 if MinZero is false", func() {
				defer func() {
					So(recover(), ShouldBeNil)
				}()

				label := "min"
				model := float64(-5)

				(Float{
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
							"Min":   float64(0),
						},
					})
				}()

				label := "min"
				model := float64(-0.1)

				(Float{
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
				model := float64(0.1)

				(Float{
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
							"Min":   float64(5),
						},
					})
				}()

				label := "min"
				model := float64(4.9)

				(Float{
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
				model := float64(5.1)

				(Float{
					Label: label,
					Model: &model,
					Min:   5,
				}).validateMin()
			})

		})

		Convey("validateMax", func() {

			Convey("Don't validate 0 if MaxZero is false", func() {
				defer func() {
					So(recover(), ShouldBeNil)
				}()

				label := "max"
				model := float64(-5)

				(Float{
					Label: label,
					Model: &model,
				}).validateMax()
			})

			Convey("Validate 0 if MaxZero is true, should panic because it's more than 0", func() {
				defer func() {
					So(recover(), ShouldResemble, &form.ErrorValidateModel{
						Key: form.LANG_NUMBER_MAX,
						Value: map[string]interface{}{
							"Label": "max",
							"Max":   float64(0),
						},
					})
				}()

				label := "max"
				model := float64(0.1)

				(Float{
					Label:   label,
					Model:   &model,
					MaxZero: true,
				}).validateMax()
			})

			Convey("Validate 0 if MaxZero is true, shoud not panic because it's less than 0", func() {
				defer func() {
					So(recover(), ShouldBeNil)
				}()

				label := "max"
				model := float64(-0.1)

				(Float{
					Label:   label,
					Model:   &model,
					MaxZero: true,
				}).validateMax()
			})

			Convey("Validate 5, should panic because it's more than 5", func() {
				defer func() {
					So(recover(), ShouldResemble, &form.ErrorValidateModel{
						Key: form.LANG_NUMBER_MAX,
						Value: map[string]interface{}{
							"Label": "max",
							"Max":   float64(5),
						},
					})
				}()

				label := "max"
				model := float64(5.1)

				(Float{
					Label: label,
					Model: &model,
					Max:   5,
				}).validateMax()
			})

			Convey("Validate 5, should not panic because it's less than 5", func() {
				defer func() {
					So(recover(), ShouldBeNil)
				}()

				label := "min"
				model := float64(4.9)

				(Float{
					Label: label,
					Model: &model,
					Max:   5,
				}).validateMax()
			})

		})

		Convey("validateStep", func() {

			Convey("Do nothing because step is set to Zero", func() {
				defer func() {
					So(recover(), ShouldBeNil)
				}()

				model := float64(3)

				(Float{Model: &model}).validateStep()
			})

			Convey("Should because panic because model is not in step", func() {
				defer func() {
					So(recover(), ShouldResemble, &form.ErrorValidateModel{
						Key: form.LANG_NUMBER_STEP,
						Value: map[string]interface{}{
							"Label": "step",
							"Step":  float64(2),
						},
					})
				}()

				label := "step"
				model := float64(3)

				(Float{
					Label: label,
					Model: &model,
					Step:  2,
				}).validateStep()
			})

			Convey("Should not panic because model is in step", func() {
				defer func() {
					So(recover(), ShouldBeNil)
				}()

				label := "step"
				model := float64(4)

				(Float{
					Label: label,
					Model: &model,
					Step:  2,
				}).validateStep()
			})

		})

		Convey("validateInList", func() {

			Convey("Should not panic, because InList is 'nil'", func() {
				defer func() {
					So(recover(), ShouldBeNil)
				}()

				(Float{}).validateInList()
			})

			Convey("Should not panic, because Model is in the List", func() {
				defer func() {
					So(recover(), ShouldBeNil)
				}()

				label := "List"
				model := float64(1.5)
				list := []float64{1.4, 1.5, 1.6}

				(Float{
					Label:  label,
					Model:  &model,
					InList: list,
				}).validateInList()
			})

			Convey("Should panic, because Model is not in the List", func() {
				defer func() {
					So(recover(), ShouldResemble, &form.ErrorValidateModel{
						Key: form.LANG_IN_LIST,
						Value: map[string]interface{}{
							"Label": "List",
							"List":  []float64{1.4, 1.5, 1.6},
						},
					})
				}()

				label := "List"
				model := float64(1.3)
				list := []float64{1.4, 1.5, 1.6}

				(Float{
					Label:  label,
					Model:  &model,
					InList: list,
				}).validateInList()
			})

		})

	})
}
