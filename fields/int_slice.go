package fields

import "encoding/json"

type IntSlice struct {
	Name     string    // Mandatory
	Label    string    // Mandatory
	Norm     *[]string // Mandatory
	Model    *[]int64  // Mandatory
	Err      *error    // Mandatory
	Required bool
	Extra    func()
}

type intSliceJson struct {
	Type     string `json:"type"`
	Name     string `json:"name"`
	Required bool   `json:"required"`
	Success  bool   `json:"success"`
	Error    string `json:"error,omitempty"`
}

func (i IntSlice) MarshalJSON() ([]byte, error) {
	return json.Marshal(intSliceJson{
		Type:     "int_slice",
		Name:     i.Name,
		Required: i.Required,
		Success:  nil == *i.Err,
		Error:    getMessageFromError(*i.Err),
	})
}
