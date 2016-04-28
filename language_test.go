package form

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestLanguage(t *testing.T) {
	langs := Langauge{
		"test":       BuildLanguageTemplate(`{{.Hello}}`),
		"listFilter": BuildLanguageTemplate(`{{.List|list "and"}}`),
		"plural":     BuildLanguageTemplate(`apple{{.Count|pluralise "s"}}`),
	}

	Convey("Bad data", t, func() {
		So(langs.Translate("test", map[string]interface{}{
			"hello": "World",
		}), ShouldEqual, "<no value>")
	})

	Convey("Good data", t, func() {
		So(langs.Translate("test", map[string]interface{}{
			"Hello": "World",
		}), ShouldEqual, "World")
	})

	Convey("Test List Filter", t, func() {
		So(langs.Translate("listFilter", map[string]interface{}{
			"List": []string{"apple", "mango", "pear"},
		}), ShouldEqual, "apple, mango and pear")
	})

	Convey("Test Plural", t, func() {
		So(langs.Translate("plural", map[string]interface{}{
			"Count": 5,
		}), ShouldEqual, "apples")
	})
}
