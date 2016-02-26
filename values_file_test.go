package form

import (
	. "github.com/smartystreets/goconvey/convey"
	"mime/multipart"
	"testing"
)

func TestValuesFile(t *testing.T) {
	Convey("Values", t, func() {

		Convey("Return empty string on 'GetOne' or 'nil' on 'GetAll'", func() {
			v := newValuesFile(&multipart.Form{
				Value: map[string][]string{},
			})

			So(v.GetOne("name"), ShouldEqual, "")
			So(v.GetAll("name"), ShouldBeNil)
		})

		Convey("Keep getting values using 'GetOne' until it's return empty string", func() {
			v := newValuesFile(&multipart.Form{
				Value: map[string][]string{
					"name": {"a", "b", "c"},
				},
			})

			So(v.GetOne("name"), ShouldEqual, "a")
			So(v.GetOne("name"), ShouldEqual, "b")
			So(v.GetOne("name"), ShouldEqual, "c")
			So(v.GetOne("name"), ShouldEqual, "")
		})

		Convey("Call 'GetAll' and compare all the values", func() {
			v := newValuesFile(&multipart.Form{
				Value: map[string][]string{
					"name": {"a", "b", "c"},
				},
			})

			So(v.GetAll("name"), ShouldResemble, []string{"a", "b", "c"})
			So(v.GetAll("name"), ShouldBeNil)
		})

		Convey("Call 'GetOne' once, than Call 'GetAll', and compare the values", func() {
			v := newValuesFile(&multipart.Form{
				Value: map[string][]string{
					"name": {"a", "b", "c"},
				},
			})

			So(v.GetOne("name"), ShouldEqual, "a")
			So(v.GetAll("name"), ShouldResemble, []string{"b", "c"})
			So(v.GetAll("name"), ShouldBeNil)
		})

	})

	Convey("Files", t, func() {

		Convey("Return 'nil' on 'GetOneFile' and 'GetAllFile'", func() {
			v := newValuesFile(&multipart.Form{
				File: map[string][]*multipart.FileHeader{},
			})

			So(v.GetOneFile("name"), ShouldBeNil)
			So(v.GetAllFile("name"), ShouldBeNil)
		})

		Convey("Keep getting file using 'GetOneFile' until it's return 'nil'", func() {
			v := newValuesFile(&multipart.Form{
				File: map[string][]*multipart.FileHeader{
					"name": {
						&multipart.FileHeader{Filename: "a"},
						&multipart.FileHeader{Filename: "b"},
						&multipart.FileHeader{Filename: "c"},
					},
				},
			})

			So(v.GetOneFile("name"), ShouldResemble, &multipart.FileHeader{Filename: "a"})
			So(v.GetOneFile("name"), ShouldResemble, &multipart.FileHeader{Filename: "b"})
			So(v.GetOneFile("name"), ShouldResemble, &multipart.FileHeader{Filename: "c"})
			So(v.GetOneFile("name"), ShouldBeNil)
		})

		Convey("Call 'GetAllFile' and compare all the files", func() {
			v := newValuesFile(&multipart.Form{
				File: map[string][]*multipart.FileHeader{
					"name": {
						&multipart.FileHeader{Filename: "a"},
						&multipart.FileHeader{Filename: "b"},
						&multipart.FileHeader{Filename: "c"},
					},
				},
			})

			So(v.GetAllFile("name"), ShouldResemble, []*multipart.FileHeader{
				&multipart.FileHeader{Filename: "a"},
				&multipart.FileHeader{Filename: "b"},
				&multipart.FileHeader{Filename: "c"},
			})
			So(v.GetAllFile("name"), ShouldBeNil)
		})

		Convey("Call 'GetOneFile' once, than Call 'GetAllFile', and compare the files", func() {
			v := newValuesFile(&multipart.Form{
				File: map[string][]*multipart.FileHeader{
					"name": {
						&multipart.FileHeader{Filename: "a"},
						&multipart.FileHeader{Filename: "b"},
						&multipart.FileHeader{Filename: "c"},
					},
				},
			})

			So(v.GetOneFile("name"), ShouldResemble, &multipart.FileHeader{Filename: "a"})
			So(v.GetAllFile("name"), ShouldResemble, []*multipart.FileHeader{
				&multipart.FileHeader{Filename: "b"},
				&multipart.FileHeader{Filename: "c"},
			})
			So(v.GetAllFile("name"), ShouldBeNil)
		})

	})

}
