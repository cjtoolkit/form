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
		So(templateListFilter("and", nil), ShouldBeEmpty)
		So(templateListFilter("and", []string{"apple"}), ShouldEqual, "apple")
		So(templateListFilter("and", []string{"apple", "mango"}), ShouldEqual, "apple and mango")
		So(templateListFilter("and", []string{"apple", "mango", "pear"}), ShouldEqual, "apple, mango and pear")
		So(templateListFilter("and", []int64{1}), ShouldEqual, "1")
		So(templateListFilter("and", []int64{1, 2}), ShouldEqual, "1 and 2")
		So(templateListFilter("and", []int64{1, 2, 3}), ShouldEqual, "1, 2 and 3")
		So(templateListFilter("and", []float64{1.5}), ShouldEqual, "1.5")
		So(templateListFilter("and", []float64{1.5, 2.5}), ShouldEqual, "1.5 and 2.5")
		So(templateListFilter("and", []float64{1.5, 2.5, 3.5}), ShouldEqual, "1.5, 2.5 and 3.5")
	})

	Convey("templatePluraliseFilter", t, func() {
		So(templatePluraliseFilter("s", 1), ShouldBeEmpty)
		So(templatePluraliseFilter("s", 5), ShouldEqual, "s")
		So(templatePluraliseFilter("s", int64(1)), ShouldBeEmpty)
		So(templatePluraliseFilter("s", int64(5)), ShouldEqual, "s")
		So(templatePluraliseFilter("s", float64(1)), ShouldBeEmpty)
		So(templatePluraliseFilter("s", float64(1.5)), ShouldEqual, "s")
		So(templatePluraliseFilter("s", uint64(1)), ShouldBeEmpty)
		So(templatePluraliseFilter("s", uint64(5)), ShouldEqual, "s")
	})
}
