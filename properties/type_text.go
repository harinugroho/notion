package properties

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Text struct {
	Content string `json:"content"`
	Link    string `json:"link"`
}

type MapText struct {
	DatabaseProperty map[string]string
	PageProperty     Text
}

func (m *MapText) Value() string {
	return m.PageProperty.Content
}

func (m *MapText) UnmarshalJSON(data []byte) error {
	var err error
	var v interface{}
	if err = json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch x := v.(type) {
	case Text:
		m.PageProperty = x
	case map[string]interface{}:
		err = json.Unmarshal(data, &m.PageProperty)
	default:
		err = fmt.Errorf("json: cannot unmarshal %v into Go value of type MapText", reflect.TypeOf(v).Name())
	}
	return err
}

func (m MapText) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.Value())
}
