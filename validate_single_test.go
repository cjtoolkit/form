package form

import (
	_ "github.com/cjtoolkit/form/lang/enGB"
	"net/http"
	"testing"
)

type testValidateSingle struct {
	Text   string
	Number int64
}

func (t *testValidateSingle) CJForm(fs *Fields) {

	// Text
	func() {
		f := fs.Init(&t.Text, "text", InputText)

		f.Mandatory()

		size := f.Size()
		size.Min = 4
		size.Max = 8
	}()

	// Number
	func() {
		f := fs.Init(&t.Number, "number", InputNumber)

		f.Mandatory()

		_range := f.RangeInt()
		_range.Min = 8
		_range.Max = 32
	}()
}

func TestValidateSingle(t *testing.T) {
	f := New(&http.Request{}, nil, "en-GB")

	var err error

	_struct := &testValidateSingle{}

	err = f.ValidateSingle(_struct, "text", []string{"hello"})

	if err != nil {
		t.Error(err.Error())
	}

	err = f.ValidateSingle(_struct, "text", []string{"hellohello"})

	if err == nil {
		t.Error("Should of exceeded max")
	}

	err = f.ValidateSingle(_struct, "text", []string{"hel"})

	if err == nil {
		t.Error("Should of gone below min")
	}

	err = f.ValidateSingle(_struct, "number", []string{"16"})

	if err != nil {
		t.Error(err.Error())
	}

	err = f.ValidateSingle(_struct, "number", []string{"33"})

	if err == nil {
		t.Error("Should of exceeded max number")
	}

	err = f.ValidateSingle(_struct, "number", []string{"7"})

	if err == nil {
		t.Error("Should of gone below min number")
	}
}
