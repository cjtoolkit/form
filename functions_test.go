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

	Convey("templateListFilter", t, func() {
		So(templateListFilter(nil, "and"), ShouldBeEmpty)
		So(templateListFilter([]string{"apple"}, "and"), ShouldEqual, "apple")
		So(templateListFilter([]string{"apple", "mango"}, "and"), ShouldEqual, "apple and mango")
		So(templateListFilter([]string{"apple", "mango", "pear"}, "and"), ShouldEqual, "apple, mango and pear")
		So(templateListFilter([]int64{1}, "and"), ShouldEqual, "1")
		So(templateListFilter([]int64{1, 2}, "and"), ShouldEqual, "1 and 2")
		So(templateListFilter([]int64{1, 2, 3}, "and"), ShouldEqual, "1, 2 and 3")
		So(templateListFilter([]float64{1.5}, "and"), ShouldEqual, "1.5")
		So(templateListFilter([]float64{1.5, 2.5}, "and"), ShouldEqual, "1.5 and 2.5")
		So(templateListFilter([]float64{1.5, 2.5, 3.5}, "and"), ShouldEqual, "1.5, 2.5 and 3.5")
	})
}
