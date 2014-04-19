package form

/* import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/gorail/core"
	"mime/multipart"
	"testing"
)

type Example struct {
	Form
	Username        string                `form:"username"`
	Password        string                `form:"password"`
	PasswordConfirm string                `form:"passwordConfirm"`
	Search          string                `form:"search"`
	Textarea        string                `form:"textarea"`
	Radio           string                `form:"radio"`
	RadioInt        int64                 `form:"radioInt"`
	RadioFloat      float64               `form:"radioFloat"`
	Select          string                `form:"select"`
	SelectInt       int64                 `form:"selectInt"`
	SelectFloat     float64               `form:"selectFloat"`
	Checkbox        bool                  `form:"checkbox"`
	Number          int64                 `form:"number"`
	NumberFloat     float64               `form:"numberFloat"`
	File            *multipart.FileHeader `form:"file"`
}

func init() {
	gob.Register(&Example{})
}

func (e *Example) UsernameType() string {
	return "input:text"
}

func (e *Example) UsernameAttr() map[string]string {
	return map[string]string{"id": "username_id", "class": "username_class"}
}

func (e *Example) UsernameRegExp() string {
	return "^([a-zA-Z]*)$"
}

func (e *Example) UsernameRegExpErr() string {
	return "Letters Only"
}

func (e *Example) UsernameMinChar() int64 {
	return 8
}

func (e *Example) UsernameMaxChar() int64 {
	return 32
}

func (e *Example) UsernameLabel() Label {
	return Label{"Username:", "username_id", nil}
}

func (e *Example) PasswordType() string {
	return "input:password"
}

func (e *Example) PasswordAttr() map[string]string {
	return map[string]string{"id": "password_id", "class": "password_class"}
}

func (e *Example) PasswordMinChar() int64 {
	return 8
}

func (e *Example) PasswordConfirmType() string {
	return "input:password"
}

func (e *Example) PasswordConfirmAttr() map[string]string {
	return map[string]string{"id": "password_confirm_id", "class": "password_confirm_class"}
}

func (e *Example) PasswordConfirmMustMatch() string {
	return "Password"
}

func (e *Example) PasswordConfirmMustMatchErr() string {
	return "Password does not match!"
}

func (e *Example) SearchType() string {
	return "input:search"
}

func (e *Example) TextareaType() string {
	return "textarea"
}

func (e *Example) TextareaRows() int64 {
	return 8
}

func (e *Example) TextareaCols() int64 {
	return 16
}

func (e *Example) RadioType() string {
	return "input:radio"
}

func (e *Example) RadioRadio() []Radio {
	return []Radio{
		{"car", "Car", false, map[string]string{"class": "car"}},
		{"motorbike", "Motorbike", true, nil},
	}
}

func (e *Example) RadioIntType() string {
	return "input:radio"
}

func (e *Example) RadioIntRadio() []RadioInt {
	return []RadioInt{
		{1, "Car", false, map[string]string{"class": "car"}},
		{2, "Motorbike", true, nil},
	}
}

func (e *Example) RadioFloatType() string {
	return "input:radio"
}

func (e *Example) RadioFloatRadio() []RadioFloat {
	return []RadioFloat{
		{1.5, "Car", false, map[string]string{"class": "car"}},
		{2.5, "Motorbike", true, nil},
	}
}

func (e *Example) SelectType() string {
	return "select"
}

func (e *Example) SelectOptions() []Option {
	return []Option{
		{
			Content: "Car",
			Value:   "car",
			Label:   "car",
		},
		{
			Content:  "Motorcycle",
			Value:    "motorcycle",
			Label:    "motorcycle",
			Selected: true,
			Attr:     map[string]string{"class": "motorcycle"},
		},
	}
}

func (e *Example) SelectIntType() string {
	return "select"
}

func (e *Example) SelectIntOptions() []OptionInt {
	return []OptionInt{
		{
			Content: "Car",
			Value:   1,
			Label:   "car",
		},
		{
			Content:  "Motorcycle",
			Value:    2,
			Label:    "motorcycle",
			Selected: true,
			Attr:     map[string]string{"class": "motorcycle"},
		},
	}
}

func (e *Example) SelectFloatType() string {
	return "select"
}

func (e *Example) SelectFloatOptions() []OptionFloat {
	return []OptionFloat{
		{
			Content: "Car",
			Value:   1.5,
			Label:   "car",
		},
		{
			Content:  "Motorcycle",
			Value:    2.5,
			Label:    "motorcycle",
			Selected: true,
			Attr:     map[string]string{"class": "motorcycle"},
		},
	}
}

func (e *Example) CheckboxType() string {
	return "input:checkbox"
}

func (e *Example) NumberType() string {
	return "input:number"
}

func (e *Example) NumberMin() int64 {
	return 5
}

func (e *Example) NumberMax() int64 {
	return 10
}

func (e *Example) NumberFloatType() string {
	return "input:number"
}

func (e *Example) NumberFloatMin() float64 {
	return 5.5
}

func (e *Example) NumberFloatMax() float64 {
	return 10.5
}

func (e *Example) FileType() string {
	return "input:file"
}

func (e *Example) FileSize() int64 {
	return 10 * 1024 * 1024
}

func (e *Example) FileAccept() []string {
	return []string{"image/jpeg", "image/png"}
}

func ExecuteExample(c *core.Context) {
	e := &Example{}
	if c.Req.Method == "GET" {
		if t, ok := c.Session().Adv().Get("form").(*Example); ok {
			e = t
			c.Session().Destroy()
		}
		fmt := c.Fmt()
		fmt.Print(`<form action="/" method="post" enctype="multipart/form-data">`)
		fmt.Println()
		Render(c.Res, e)
		fmt.Print(`<input type="submit" name="submit" value="Submit">
</form>`)
	} else if c.Req.Method == "POST" {
		if !Validate(e, c) {
			c.Session().Adv().Set("form", e)
			c.Session().Adv().Save()
			c.Url().Redirect("/")
			return
		}
		c.Fmt().Print(`Success!`)
	}
}

func TestExample(t *testing.T) {
	buf := &bytes.Buffer{}

	Render(buf, &Example{
		Username:        "Username",
		Password:        "Password",
		PasswordConfirm: "Confirm",
		Textarea:        "Textarea",
		Radio:           "car",
		Select:          "car",
		Checkbox:        true,
	})

	fmt.Println(buf.String())
} */
