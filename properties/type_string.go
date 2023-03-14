package properties

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type MapString struct {
	IsNull           bool
	DatabaseProperty map[string]string
	PageProperty     string
}

func (m *MapString) Value() string {
	return m.PageProperty
}

func (m *MapString) UnmarshalJSON(data []byte) error {
	var err error
	var v interface{}
	if err = json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch x := v.(type) {
	case string:
		m.PageProperty = x
	case map[string]interface{}:
		err = json.Unmarshal(data, &m.DatabaseProperty)
	case nil:
		m.IsNull = true
	default:
		err = fmt.Errorf("json: cannot unmarshal %v into Go value of type MapString", reflect.TypeOf(v).Name())
	}
	return err
}

func (m MapString) MarshalJSON() ([]byte, error) {
	if m.IsNull {
		return []byte("null"), nil
	}
	return json.Marshal(m.Value())
}
