package fields

import (
	"github.com/cjtoolkit/form"
	. "github.com/smartystreets/goconvey/convey"
	"regexp"
	"testing"
)

func TestString(t *testing.T) {
	Convey("PreCheck", t, func() {
		Convey("Should panic because Name is empty string", func() {
			go panicTrap(func() { (String{}).PreCheck() })

			So(<-panicChannel, ShouldEqual, form.ErrorPreCheck("Text Field: Name cannot be empty string"))
		})

		Convey("Should panic because Label is empty string", func() {
			go panicTrap(func() {
				(String{
					Name: "hello",
				}).PreCheck()
			})

			So(<-panicChannel, ShouldEqual, form.ErrorPreCheck("Text Field: hello: Label cannot be empty string"))
		})

		Convey("Should panic because Norm is nil value", func() {
			go panicTrap(func() {
				(String{
					Name:  "hello",
					Label: "Hello",
				}).PreCheck()
			})

			So(<-panicChannel, ShouldEqual, form.ErrorPreCheck("Text Field: hello: Norm cannot be nil value"))
		})

		Convey("Should panic because Model is nil value", func() {
			var norm string

			go panicTrap(func() {
				(String{
					Name:  "hello",
					Label: "Hello",
					Norm:  &norm,
				}).PreCheck()
			})

			So(<-panicChannel, ShouldEqual, form.ErrorPreCheck("Text Field: hello: Model cannot be nil value"))
		})

		Convey("Should panic because Err is nil value", func() {
			var norm, model string

			go panicTrap(func() {
				(String{
					Name:  "hello",
					Label: "Hello",
					Norm:  &norm,
					Model: &model,
				}).PreCheck()
			})

			So(<-panicChannel, ShouldEqual, form.ErrorPreCheck("Text Field: hello: Err cannot be nil value"))
		})

		Convey("Every mandatory field is in order, so therefore should not panic", func() {
			var norm, model string
			var err error

			go panicTrap(func() {
				(String{
					Name:  "hello",
					Label: "Hello",
					Norm:  &norm,
					Model: &model,
					Err:   &err,
				}).PreCheck()
			})

			So(<-panicChannel, ShouldBeNil)
		})
	})

	Convey("ValidateModel", t, func() {

		Convey("validateRequired", func() {

			Convey("Should not panic because field is not required", func() {
				go panicTrap(func() { (String{}).validateRequired() })

				So(<-panicChannel, ShouldBeNil)
			})

			Convey("Should panic because required field is an empty string", func() {
				model := ""

				go panicTrap(func() {
					(String{
						Label:    "test",
						Model:    &model,
						Required: true,
					}).validateRequired()
				})

				So(<-panicChannel, ShouldResemble, &form.ErrorValidateModel{
					Key: form.LANG_FIELD_REQUIRED,
					Value: map[string]interface{}{
						"Label": "test",
					},
				})
			})

			Convey("Should not panic because required field is not an empty string", func() {
				model := "hello"

				go panicTrap(func() {
					(String{
						Label:    "test",
						Model:    &model,
						Required: true,
					}).validateRequired()
				})

				So(<-panicChannel, ShouldBeNil)
			})

		})

		Convey("validateMinRune", func() {

			Convey("Should not panic because MinChar has not been populated", func() {
				go panicTrap(func() { (String{}).validateMinRune() })

				So(<-panicChannel, ShouldBeNil)
			})

			Convey("Should panic because model is less than MinChar", func() {
				model := "he"

				go panicTrap(func() {
					(String{
						Label:   "test",
						MinRune: 4,
						Model:   &model,
					}).validateMinRune()
				})

				So(<-panicChannel, ShouldResemble, &form.ErrorValidateModel{
					Key: form.LANG_MIN_CHAR,
					Value: map[string]interface{}{
						"Label":   "test",
						"MinRune": 4,
					},
				})
			})

			Convey("Should not panic because model is more than MinChar", func() {
				model := "hello"

				go panicTrap(func() {
					(String{
						Label:   "test",
						MinRune: 4,
						Model:   &model,
					}).validateMinRune()
				})

				So(<-panicChannel, ShouldBeNil)
			})

		})

		Convey("validateMaxRune", func() {

			Convey("Should not panic because MaxChar has not been populated", func() {
				go panicTrap(func() { (String{}).validateMaxRune() })

				So(<-panicChannel, ShouldBeNil)
			})

			Convey("Should panic because model is greater than MaxChar", func() {
				model := "hello"

				go panicTrap(func() {
					(String{
						Label:   "test",
						MaxRune: 4,
						Model:   &model,
					}).validateMaxRune()
				})

				So(<-panicChannel, ShouldResemble, &form.ErrorValidateModel{
					Key: form.LANG_MAX_CHAR,
					Value: map[string]interface{}{
						"Label":   "test",
						"MaxRune": 4,
					},
				})
			})

			Convey("Should not panic because model is less tahn MaxChar", func() {
				model := "he"

				go panicTrap(func() {
					(String{
						Label:   "test",
						MaxRune: 4,
						Model:   &model,
					}).validateMaxRune()
				})

				So(<-panicChannel, ShouldBeNil)
			})

		})

		Convey("validateMustMatch", func() {

			Convey("Should not panic because MustMatchModel and/or MustMatchLabel has been populated", func() {
				mustMatchModel := "test"
				mustMatchLabel := "Test"

				go panicTrap(func() { (String{}).validateMustMatch() })
				go panicTrap(func() {
					(String{
						MustMatchModel: &mustMatchModel,
					}).validateMustMatch()
				})
				go panicTrap(func() {
					(String{
						MustMatchLabel: mustMatchLabel,
					}).validateMustMatch()
				})

				So(<-panicChannel, ShouldBeNil)
				So(<-panicChannel, ShouldBeNil)
				So(<-panicChannel, ShouldBeNil)
			})

			Convey("Should not panic because MustMatchModel and Model are identical to each other", func() {
				model := "apple"
				mustMatchModel := "apple"
				mustMatchLabel := "Matching Fruit"

				go panicTrap(func() {
					(String{
						Model:          &model,
						MustMatchModel: &mustMatchModel,
						MustMatchLabel: mustMatchLabel,
					}).validateMustMatch()
				})

				So(<-panicChannel, ShouldBeNil)
			})

			Convey("Should panic because MustMatchModel and Model are not identical to each other", func() {
				model := "apple"
				label := "Fruit"
				mustMatchModel := "orange"
				mustMatchLabel := "Matching Fruit"

				go panicTrap(func() {
					(String{
						Label:          label,
						Model:          &model,
						MustMatchLabel: mustMatchLabel,
						MustMatchModel: &mustMatchModel,
					}).validateMustMatch()
				})

				So(<-panicChannel, ShouldResemble, &form.ErrorValidateModel{
					Key: form.LANG_MUST_MATCH,
					Value: map[string]interface{}{
						"Label":          "Fruit",
						"MustMatchLabel": "Matching Fruit",
					},
				})
			})

		})

		Convey("validatePattern", func() {

			Convey("Should not panic, because Pattern is 'nil'", func() {
				go panicTrap(func() { (String{}).validatePattern() })

				So(<-panicChannel, ShouldBeNil)
			})

			Convey("Should not panic, because Model matches Pattern", func() {
				label := "Pattern"
				pattern := regexp.MustCompile(`\d`)
				model := "5"

				go panicTrap(func() {
					(String{
						Label:   label,
						Pattern: pattern,
						Model:   &model,
					}).validatePattern()
				})

				So(<-panicChannel, ShouldBeNil)
			})

			Convey("Panic, because Model does not match Pattern", func() {
				label := "Pattern"
				pattern := regexp.MustCompile(`\d`)
				model := "a"

				go panicTrap(func() {
					(String{
						Label:   label,
						Pattern: pattern,
						Model:   &model,
					}).validatePattern()
				})

				So(<-panicChannel, ShouldResemble, &form.ErrorValidateModel{
					Key: form.LANG_PATTERN,
					Value: map[string]interface{}{
						"Label":   "Pattern",
						"Pattern": `\d`,
					},
				})
			})

		})

		Convey("validateInList", func() {

			Convey("Should not panic, because InList is 'nil'", func() {
				go panicTrap(func() { (String{}).validateInList() })

				So(<-panicChannel, ShouldBeNil)
			})

			Convey("Should not panic, because Model is in the List", func() {
				label := "List"
				model := "apple"
				list := []string{"orange", "apple", "pear"}

				go panicTrap(func() {
					(String{
						Label:  label,
						Model:  &model,
						InList: list,
					}).validateInList()
				})

				So(<-panicChannel, ShouldBeNil)
			})

			Convey("Should panic, because Model is not in the List", func() {
				label := "List"
				model := "mango"
				list := []string{"orange", "apple", "pear"}

				go panicTrap(func() {
					(String{
						Label:  label,
						Model:  &model,
						InList: list,
					}).validateInList()
				})

				So(<-panicChannel, ShouldResemble, &form.ErrorValidateModel{
					Key: form.LANG_IN_LIST,
					Value: map[string]interface{}{
						"Label": "List",
						"List":  []string{"orange", "apple", "pear"},
					},
				})
			})

		})

	})
}
