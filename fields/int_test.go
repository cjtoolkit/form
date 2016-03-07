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
}
