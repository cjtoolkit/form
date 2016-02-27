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
}
