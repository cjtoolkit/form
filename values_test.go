package form

import (
	. "github.com/smartystreets/goconvey/convey"
	"net/url"
	"testing"
)

func TestValues(t *testing.T) {
	Convey("Return empty string on 'GetOne' or 'nil' on 'GetAll'", t, func() {
		v := newValues(url.Values{})

		So(v.GetOne("name"), ShouldEqual, "")
		So(v.GetAll("name"), ShouldBeNil)
	})

	Convey("Keep getting values using 'GetOne' until it's return empty string", t, func() {
		v := newValues(url.Values{
			"name": {"a", "b", "c"},
		})

		So(v.GetOne("name"), ShouldEqual, "a")
		So(v.GetOne("name"), ShouldEqual, "b")
		So(v.GetOne("name"), ShouldEqual, "c")
		So(v.GetOne("name"), ShouldEqual, "")
	})

	Convey("Call 'GetAll' and compare all the values", t, func() {
		v := newValues(url.Values{
			"name": {"a", "b", "c"},
		})

		So(v.GetAll("name"), ShouldResemble, []string{"a", "b", "c"})
		So(v.GetAll("name"), ShouldBeNil)
	})

	Convey("Call 'GetOne' once, than Call 'GetAll', and compared the values", t, func() {
		v := newValues(url.Values{
			"name": {"a", "b", "c"},
		})

		So(v.GetOne("name"), ShouldEqual, "a")
		So(v.GetAll("name"), ShouldResemble, []string{"b", "c"})
		So(v.GetAll("name"), ShouldBeNil)
	})
}
