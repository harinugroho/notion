package properties

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type External struct {
	Url string `json:"url"`
}

type MapExternal struct {
	DatabaseProperty map[string]string
	PageProperty     External
}

func (m *MapExternal) Value() string {
	return m.PageProperty.Url
}

func (m *MapExternal) UnmarshalJSON(data []byte) error {
	var err error
	var v interface{}
	if err = json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch x := v.(type) {
	case External:
		m.PageProperty = x
	case map[string]interface{}:
		err = json.Unmarshal(data, &m.PageProperty)
	default:
		err = fmt.Errorf("json: cannot unmarshal %v into Go value of type NullFile", reflect.TypeOf(v).Name())
	}
	return err
}

func (m MapExternal) MarshalJSON() ([]byte, error) {
	if m.Value() == "" {
		return []byte("null"), nil
	}
	return json.Marshal(m.Value())
}
