package form

import (
	. "github.com/smartystreets/goconvey/convey"
	"mime/multipart"
	"testing"
)

func TestFunctions(t *testing.T) {
	Convey("GetOneFile", t, func() {

		Convey("Returns nil if 'values' does not implement 'ValuesFileInterface'", func() {
			So(GetOneFile(nil, "name"), ShouldBeNil)
		})

		Convey("Returns 'FileHeader' if values does implement 'ValuesFileInterface'", func() {
			values := newValuesFile(&multipart.Form{
				File: map[string][]*multipart.FileHeader{
					"name": {
						&multipart.FileHeader{Filename: "a"},
					},
				},
			})

			So(GetOneFile(values, "name"), ShouldResemble, &multipart.FileHeader{Filename: "a"})
		})

	})

	Convey("GetAllFile", t, func() {

		Convey("Returns nil if 'values' does not implement 'ValuesFileInterface'", func() {
			So(GetAllFile(nil, "name"), ShouldBeNil)
		})

		Convey("Returns Multiple 'FileHeader' if values does implement 'ValuesFileInterface'", func() {
			values := newValuesFile(&multipart.Form{
				File: map[string][]*multipart.FileHeader{
					"name": {
						&multipart.FileHeader{Filename: "a"},
						&multipart.FileHeader{Filename: "b"},
						&multipart.FileHeader{Filename: "c"},
					},
				},
			})

			So(GetAllFile(values, "name"), ShouldResemble, []*multipart.FileHeader{
				&multipart.FileHeader{Filename: "a"},
				&multipart.FileHeader{Filename: "b"},
				&multipart.FileHeader{Filename: "c"},
			})
		})

	})

	Convey("BuildLanguageTemplate", t, func() {

		Convey("Panic if template builder fail parse properly", func() {
			defer func() {
				So(recover(), ShouldNotBeNil)
			}()

			BuildLanguageTemplate("{{.Hello")
		})

		Convey("Build build template properly, therefore does not panic", func() {
			defer func() {
				So(recover(), ShouldBeNil)
			}()

			BuildLanguageTemplate("{{.Hello}}")
		})

	})
}
