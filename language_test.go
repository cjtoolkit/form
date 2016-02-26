package form

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestLanguage(t *testing.T) {
	// start let

	langs := Langauge{
		"test": BuildLanguageTemplate("{{.Hello}}"),
	}

	// end let

	Convey("Bad parameters", t, func() {
		So(langs.Translate("test", map[string]interface{}{
			"hello": "World",
		}), ShouldEqual, "<no value>")
	})

	Convey("Good parameters", t, func() {
		So(langs.Translate("test", map[string]interface{}{
			"Hello": "World",
		}), ShouldEqual, "World")
	})
}
