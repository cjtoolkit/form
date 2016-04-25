package fields

import (
	"bytes"
	"fmt"
	"github.com/cjtoolkit/form"
	. "github.com/smartystreets/goconvey/convey"
	"mime/multipart"
	"net/textproto"
	"testing"
)

var multipartForm *multipart.Form

func init() {
	buf := &bytes.Buffer{}

	mw := multipart.NewWriter(buf)

	h := make(textproto.MIMEHeader)

	h.Set("Content-Disposition", `form-data; name="text4"; filename="text4.txt"`)
	h.Set("Content-Type", "text/plain")

	w, _ := mw.CreatePart(h)

	fmt.Fprint(w, "aaaa")

	h = make(textproto.MIMEHeader)

	h.Set("Content-Disposition", `form-data; name="text8"; filename="text8.txt"`)
	h.Set("Content-Type", "text/plain")

	w, _ = mw.CreatePart(h)

	fmt.Fprint(w, "aaaaaaaa")

	h = make(textproto.MIMEHeader)

	h.Set("Content-Disposition", `form-data; name="img"; filename="img.jpg"`)
	h.Set("Content-Type", "image/jpeg")

	w, _ = mw.CreatePart(h)

	fmt.Fprint(w, "iiii")

	boundary := mw.Boundary()

	mw.Close()

	multipartForm, _ = multipart.NewReader(buf, boundary).ReadForm(1 * 1024 * 1024)
}

func TestFile(t *testing.T) {
	Convey("PreCheck", t, func() {
		Convey("Panic because name is empty", func() {
			go panicTrap(func() {
				(File{}).PreCheck()
			})

			So(<-panicChannel, ShouldEqual, form.ErrorPreCheck("File Field: Name cannot be empty string"))
		})

		Convey("Panic because label is empty", func() {
			go panicTrap(func() {
				(File{
					Name: "test",
				}).PreCheck()
			})

			So(<-panicChannel, ShouldEqual, form.ErrorPreCheck("File Field: test: Label cannot be empty string"))
		})

		Convey("Panic because File is nil", func() {
			go panicTrap(func() {
				(File{
					Name:  "test",
					Label: "test",
				}).PreCheck()
			})

			So(<-panicChannel, ShouldEqual, form.ErrorPreCheck("File Field: test: File cannot be nil value"))
		})

		Convey("Panic because Err is nil", func() {
			var file *multipart.FileHeader

			go panicTrap(func() {
				(File{
					Name:  "test",
					Label: "test",
					File:  &file,
				}).PreCheck()
			})

			So(<-panicChannel, ShouldEqual, form.ErrorPreCheck("File Field: test: Err cannot be nil value"))
		})

		Convey("Everything checks out", func() {
			var file *multipart.FileHeader
			var err error

			go panicTrap(func() {
				(File{
					Name:  "test",
					Label: "test",
					File:  &file,
					Err:   &err,
				}).PreCheck()
			})

			So(<-panicChannel, ShouldBeNil)
		})
	})

	Convey("ValidateModel", t, func() {

		Convey("validateRequired", func() {

			Convey("Do nothing because it's not required", func() {
				var file *multipart.FileHeader

				go panicTrap(func() {
					(File{
						File: &file,
					}).validateRequired()
				})

				So(<-panicChannel, ShouldBeNil)
			})

			Convey("Panic because file is nil while is required", func() {
				var file *multipart.FileHeader

				go panicTrap(func() {
					(File{
						File:     &file,
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

			Convey("Should not panic because file is not nil while is required", func() {
				var file *multipart.FileHeader = multipartForm.File["text4"][0]

				go panicTrap(func() {
					(File{
						File:     &file,
						Required: true,
					}).validateRequired()
				})

				So(<-panicChannel, ShouldBeNil)
			})

		})

		Convey("validateMime", func() {

			Convey("Does nothing because MIME has not been set", func() {
				var file *multipart.FileHeader = multipartForm.File["text4"][0]

				go panicTrap(func() {
					(File{
						File: &file,
					}).validateMime()
				})

				So(<-panicChannel, ShouldBeNil)
			})

			Convey("Panic because file is not image/jpeg", func() {
				var file *multipart.FileHeader = multipartForm.File["text4"][0]

				go panicTrap(func() {
					(File{
						File: &file,
						Mime: []string{"image/jpeg"},
					}).validateMime()
				})

				So(<-panicChannel, ShouldResemble, &form.ErrorValidateModel{
					Key: form.LANG_FILE_MIME,
					Value: map[string]interface{}{
						"Label": "",
						"Mime":  []string{"image/jpeg"},
					},
				})
			})

			Convey("Should not panic because file is image/jpeg", func() {
				var file *multipart.FileHeader = multipartForm.File["img"][0]

				go panicTrap(func() {
					(File{
						File: &file,
						Mime: []string{"image/jpeg"},
					}).validateMime()
				})

				So(<-panicChannel, ShouldBeNil)
			})

		})

		Convey("validateSizeInByte", func() {

			Convey("Does nothing because SizeInByte has not been set", func() {
				var file *multipart.FileHeader = multipartForm.File["text8"][0]

				go panicTrap(func() {
					(File{
						File: &file,
					}).validateSizeInByte()
				})

				So(<-panicChannel, ShouldBeNil)
			})

			Convey("Panic because file size is greater than 6", func() {
				var file *multipart.FileHeader = multipartForm.File["text8"][0]

				go panicTrap(func() {
					(File{
						File:       &file,
						SizeInByte: 6,
					}).validateSizeInByte()
				})

				So(<-panicChannel, ShouldResemble, &form.ErrorValidateModel{
					Key: form.LANG_FILE_SIZE,
					Value: map[string]interface{}{
						"Label": "",
						"Size":  int64(6),
					},
				})
			})

			Convey("Should not panic because file size is less than 6", func() {
				var file *multipart.FileHeader = multipartForm.File["text4"][0]

				go panicTrap(func() {
					(File{
						File:       &file,
						SizeInByte: 6,
					}).validateSizeInByte()
				})

				So(<-panicChannel, ShouldBeNil)
			})

		})

	})
}
