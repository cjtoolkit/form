package fields

import (
	"github.com/cjtoolkit/form"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	form.FormFieldInterfaceCheck(Time{})

	Convey("PreCheck", t, func() {

		Convey("Should panic because name is empty", func() {
			go panicTrap(func() {
				(Time{}).PreCheck()
			})

			So(<-panicChannel, ShouldEqual, form.ErrorPreCheck("Time Field: Name cannot be empty string"))
		})

		Convey("Should panic because label is emtpy", func() {
			go panicTrap(func() {
				(Time{
					Name: "test",
				}).PreCheck()
			})

			So(<-panicChannel, ShouldEqual,
				form.ErrorPreCheck("Time Field: test: Label cannot be empty string"))
		})

		Convey("Should panic because norm is nil", func() {
			go panicTrap(func() {
				(Time{
					Name:  "test",
					Label: "test",
				}).PreCheck()
			})

			So(<-panicChannel, ShouldEqual,
				form.ErrorPreCheck("Time Field: test: Norm cannot be nil value"))
		})

		Convey("Should panic because model is nil", func() {
			var norm string

			go panicTrap(func() {
				(Time{
					Name:  "test",
					Label: "test",
					Norm:  &norm,
				}).PreCheck()
			})

			So(<-panicChannel, ShouldEqual,
				form.ErrorPreCheck("Time Field: test: Model cannot be nil value"))
		})

		Convey("Should panic because err is nil", func() {
			var norm string
			var model time.Time

			go panicTrap(func() {
				(Time{
					Name:  "test",
					Label: "test",
					Norm:  &norm,
					Model: &model,
				}).PreCheck()
			})

			So(<-panicChannel, ShouldEqual,
				form.ErrorPreCheck("Time Field: test: Err cannot be nil value"))
		})

		Convey("Should panic because location is nil", func() {
			var norm string
			var model time.Time
			var err error

			go panicTrap(func() {
				(Time{
					Name:  "test",
					Label: "test",
					Norm:  &norm,
					Model: &model,
					Err:   &err,
				}).PreCheck()
			})

			So(<-panicChannel, ShouldEqual,
				form.ErrorPreCheck("Time Field: test: Location cannot be nil value"))
		})

		Convey("Should panic because formats is nil or 0", func() {
			var norm string
			var model time.Time
			var err error

			go panicTrap(func() {
				(Time{
					Name:     "test",
					Label:    "test",
					Norm:     &norm,
					Model:    &model,
					Err:      &err,
					Location: time.UTC,
				}).PreCheck()
			})

			So(<-panicChannel, ShouldEqual,
				form.ErrorPreCheck("Time Field: test: Formats cannot be nil value or empty"))

			go panicTrap(func() {
				(Time{
					Name:     "test",
					Label:    "test",
					Norm:     &norm,
					Model:    &model,
					Err:      &err,
					Location: time.UTC,
					Formats:  []string{},
				}).PreCheck()
			})

			So(<-panicChannel, ShouldEqual,
				form.ErrorPreCheck("Time Field: test: Formats cannot be nil value or empty"))
		})

		Convey("Everything checks out", func() {
			var norm string
			var model time.Time
			var err error

			go panicTrap(func() {
				(Time{
					Name:     "test",
					Label:    "test",
					Norm:     &norm,
					Model:    &model,
					Err:      &err,
					Location: time.UTC,
					Formats:  []string{"hi"},
				}).PreCheck()
			})

			So(<-panicChannel, ShouldBeNil)
		})

	})

	Convey("ReverseTransform", t, func() {

		Convey("Should panic because it failed to parse time", func() {
			var model time.Time
			norm := "abc"

			go panicTrap(func() {
				(Time{
					Model:    &model,
					Norm:     &norm,
					Formats:  TimeFormats(),
					Location: time.UTC,
				}).ReverseTransform()
			})

			So(<-panicChannel, ShouldResemble, form.ErrorReverseTransform{
				Key: form.LANG_TIME_FORMAT,
				Value: Time{
					Model:    &model,
					Norm:     &norm,
					Formats:  TimeFormats(),
					Location: time.UTC,
				},
			})
		})

		Convey("Should not panic because it parsed time", func() {
			var model time.Time
			norm := "19:30"

			go panicTrap(func() {
				(Time{
					Model:    &model,
					Norm:     &norm,
					Formats:  TimeFormats(),
					Location: time.UTC,
				}).ReverseTransform()
			})

			So(<-panicChannel, ShouldBeNil)
			So(model.Hour(), ShouldEqual, 19)
			So(model.Minute(), ShouldEqual, 30)
		})

	})

	Convey("ValidateModel", t, func() {

		Convey("validateRequired", func() {

			Convey("Do nothing because it's not required", func() {
				go panicTrap(func() {
					(Time{}).validateRequired()
				})

				So(<-panicChannel, ShouldBeNil)
			})

			Convey("Panic because it is required", func() {
				var model time.Time

				go panicTrap(func() {
					(Time{
						Model:    &model,
						Required: true,
					}).validateRequired()
				})

				So(<-panicChannel, ShouldResemble, &form.ErrorValidateModel{
					Key: form.LANG_FIELD_REQUIRED,
					Value: Time{
						Model:    &model,
						Required: true,
					},
				})
			})

			Convey("Should not panic because it's there as required", func() {
				var model time.Time = time.Now()

				go panicTrap(func() {
					(Time{
						Model:    &model,
						Required: true,
					}).validateRequired()
				})

				So(<-panicChannel, ShouldBeNil)
			})

		})

		Convey("validateMin", func() {
			Convey("Do nothing as min and minzero not been specified", func() {
				go panicTrap(func() {
					(Time{}).validateMin()
				})

				So(<-panicChannel, ShouldBeNil)
			})

			Convey("Panic because it's less than minimum", func() {
				model := time.Unix(3, 0).In(time.UTC)

				go panicTrap(func() {
					(Time{
						Model:    &model,
						Location: time.UTC,
						Formats:  []string{"15:04:05"},
						Min:      time.Unix(4, 0).In(time.UTC),
					}).validateMin()
				})

				So(<-panicChannel, ShouldResemble, &form.ErrorValidateModel{
					Key: form.LANG_TIME_MIN,
					Value: Time{
						Model:    &model,
						Location: time.UTC,
						Formats:  []string{"15:04:05"},
						Min:      time.Unix(4, 0).In(time.UTC),
					},
				})
			})

			Convey("Should not because it's more than minimum", func() {
				model := time.Unix(5, 0).In(time.UTC)

				go panicTrap(func() {
					(Time{
						Model:    &model,
						Location: time.UTC,
						Formats:  []string{"15:04:05"},
						Min:      time.Unix(4, 0).In(time.UTC),
					}).validateMin()
				})

				So(<-panicChannel, ShouldBeNil)
			})
		})

		Convey("validateMax", func() {
			Convey("Do nothing as max and maxzero not been specified", func() {
				go panicTrap(func() {
					(Time{}).validateMax()
				})

				So(<-panicChannel, ShouldBeNil)
			})

			Convey("Panic because it's more than maximum", func() {
				model := time.Unix(5, 0).In(time.UTC)

				go panicTrap(func() {
					(Time{
						Model:    &model,
						Location: time.UTC,
						Formats:  []string{"15:04:05"},
						Max:      time.Unix(4, 0).In(time.UTC),
					}).validateMax()
				})

				So(<-panicChannel, ShouldResemble, &form.ErrorValidateModel{
					Key: form.LANG_TIME_MAX,
					Value: Time{
						Model:    &model,
						Location: time.UTC,
						Formats:  []string{"15:04:05"},
						Max:      time.Unix(4, 0).In(time.UTC),
					},
				})
			})

			Convey("Should not because it's less than maximum", func() {
				model := time.Unix(3, 0).In(time.UTC)

				go panicTrap(func() {
					(Time{
						Model:    &model,
						Location: time.UTC,
						Formats:  []string{"15:04:05"},
						Max:      time.Unix(4, 0).In(time.UTC),
					}).validateMax()
				})

				So(<-panicChannel, ShouldBeNil)
			})
		})

	})
}
