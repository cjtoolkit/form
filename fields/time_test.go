package fields

import (
	"github.com/cjtoolkit/form"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
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
					Location: time.Local,
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
					Location: time.Local,
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
					Location: time.Local,
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
					Model: &model,
					Norm:  &norm,
					Formats: TimeFormats(),
					Location: time.Local,
				}).ReverseTransform()
			})

			So(<-panicChannel, ShouldResemble, form.ErrorReverseTransform{
				Key: form.LANG_TIME_FORMAT,
				Value: map[string]interface{}{
					"Label": "",
				},
			})
		})

		Convey("Should not panic because it parsed time", func() {
			var model time.Time
			norm := "19:30"

			go panicTrap(func() {
				(Time{
					Model: &model,
					Norm:  &norm,
					Formats: TimeFormats(),
					Location: time.Local,
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
					Value: map[string]interface{}{
						"Label": "",
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

	})
}
