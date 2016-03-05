package fields

import (
	"github.com/cjtoolkit/form"
	. "github.com/smartystreets/goconvey/convey"
	"regexp"
	"testing"
)

func TestText(t *testing.T) {
	Convey("PreCheck", t, func() {
		Convey("Should panic because Name is empty string", func() {
			defer func() {
				So(recover(), ShouldEqual, form.ErrorPreCheck("Text Field: Name cannot be empty string"))
			}()

			(Text{}).PreCheck()
		})

		Convey("Should panic because Label is empty string", func() {
			defer func() {
				So(recover(), ShouldEqual, form.ErrorPreCheck("Text Field: hello: Label cannot be empty string"))
			}()

			(Text{
				Name: "hello",
			}).PreCheck()
		})

		Convey("Should panic because Norm is nil value", func() {
			defer func() {
				So(recover(), ShouldEqual, form.ErrorPreCheck("Text Field: hello: Norm cannot be nil value"))
			}()

			(Text{
				Name:  "hello",
				Label: "Hello",
			}).PreCheck()
		})

		Convey("Should panic because Model is nil value", func() {
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

		Convey("Should panic because Err is nil value", func() {
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

	Convey("ValidateModel", t, func() {

		Convey("validateRequired", func() {

			Convey("Should not panic because field is not required", func() {
				defer func() {
					So(recover(), ShouldBeNil)
				}()

				(Text{}).validateRequired()
			})

			Convey("Should panic because required field is an empty string", func() {
				defer func() {
					So(recover(), ShouldResemble, &form.ErrorValidateModel{
						Key: form.LANG_FIELD_REQUIRED,
						Value: map[string]interface{}{
							"Label": "test",
						},
					})
				}()

				model := ""

				(Text{
					Label:    "test",
					Model:    &model,
					Required: true,
				}).validateRequired()
			})

			Convey("Should not panic because required field is not an empty string", func() {
				defer func() {
					So(recover(), ShouldBeNil)
				}()

				model := "hello"

				(Text{
					Label:    "test",
					Model:    &model,
					Required: true,
				}).validateRequired()
			})

		})

		Convey("validateMinChar", func() {

			Convey("Should not panic because MinChar has not been populated", func() {
				defer func() {
					So(recover(), ShouldBeNil)
				}()

				(Text{}).validateMinChar()
			})

			Convey("Should panic because model is less than MinChar", func() {
				defer func() {
					So(recover(), ShouldResemble, &form.ErrorValidateModel{
						Key: form.LANG_MIN_CHAR,
						Value: map[string]interface{}{
							"Label":   "test",
							"MinChar": 4,
						},
					})
				}()

				model := "he"

				(Text{
					Label:   "test",
					MinChar: 4,
					Model:   &model,
				}).validateMinChar()
			})

			Convey("Should not panic because model is more than MinChar", func() {
				defer func() {
					So(recover(), ShouldBeNil)
				}()

				model := "hello"

				(Text{
					Label:   "test",
					MinChar: 4,
					Model:   &model,
				}).validateMinChar()
			})

		})

		Convey("validateMaxChar", func() {

			Convey("Should not panic because MaxChar has not been populated", func() {
				defer func() {
					So(recover(), ShouldBeNil)
				}()

				(Text{}).validateMaxChar()
			})

			Convey("Should panic because model is greater than MaxChar", func() {
				defer func() {
					So(recover(), ShouldResemble, &form.ErrorValidateModel{
						Key: form.LANG_MAX_CHAR,
						Value: map[string]interface{}{
							"Label":   "test",
							"MaxChar": 4,
						},
					})
				}()

				model := "hello"

				(Text{
					Label:   "test",
					MaxChar: 4,
					Model:   &model,
				}).validateMaxChar()
			})

			Convey("Should not panic because model is less tahn MinChar", func() {
				defer func() {
					So(recover(), ShouldBeNil)
				}()

				model := "he"

				(Text{
					Label:   "test",
					MaxChar: 4,
					Model:   &model,
				}).validateMaxChar()
			})

		})

		Convey("validateMustMatch", func() {

			Convey("Should not panic because MustMatchModel and/or MustMatchLabel has been populated", func() {
				defer func() {
					So(recover(), ShouldBeNil)
				}()

				mustMatchModel := "test"
				mustMatchLabel := "Test"

				(Text{}).validateMustMatch()
				(Text{
					MustMatchModel: &mustMatchModel,
				}).validateMustMatch()
				(Text{
					MustMatchLabel: mustMatchLabel,
				}).validateMustMatch()
			})

			Convey("Should not panic because MustMatchModel and Model are identical to each other", func() {
				defer func() {
					So(recover(), ShouldBeNil)
				}()

				model := "apple"
				mustMatchModel := "apple"
				mustMatchLabel := "Matching Fruit"

				(Text{
					Model:          &model,
					MustMatchModel: &mustMatchModel,
					MustMatchLabel: mustMatchLabel,
				}).validateMustMatch()
			})

			Convey("Should panic because MustMatchModel and Model are not identical to each other", func() {
				defer func() {
					So(recover(), ShouldResemble, &form.ErrorValidateModel{
						Key: form.LANG_MUST_MATCH,
						Value: map[string]interface{}{
							"Label":          "Fruit",
							"MustMatchLabel": "Matching Fruit",
						},
					})
				}()

				model := "apple"
				label := "Fruit"
				mustMatchModel := "orange"
				mustMatchLabel := "Matching Fruit"

				(Text{
					Label:          label,
					Model:          &model,
					MustMatchLabel: mustMatchLabel,
					MustMatchModel: &mustMatchModel,
				}).validateMustMatch()
			})

		})

		Convey("validatePattern", func() {

			Convey("Should not panic, because Pattern is 'nil'", func() {
				defer func() {
					So(recover(), ShouldBeNil)
				}()

				(Text{}).validatePattern()
			})

			Convey("Should not panic, because Model matches Pattern", func() {
				defer func() {
					So(recover(), ShouldBeNil)
				}()

				label := "Pattern"
				pattern := regexp.MustCompile(`\d`)
				model := "5"

				(Text{
					Label:   label,
					Pattern: pattern,
					Model:   &model,
				}).validatePattern()
			})

			Convey("Panic, because Model does not match Pattern", func() {
				defer func() {
					So(recover(), ShouldResemble, &form.ErrorValidateModel{
						Key: form.LANG_PATTERN,
						Value: map[string]interface{}{
							"Label":   "Pattern",
							"Pattern": `\d`,
						},
					})
				}()

				label := "Pattern"
				pattern := regexp.MustCompile(`\d`)
				model := "a"

				(Text{
					Label:   label,
					Pattern: pattern,
					Model:   &model,
				}).validatePattern()
			})

		})

	})
}
