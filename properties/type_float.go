package properties

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type MapFloat struct {
	IsNull           bool
	DatabaseProperty map[string]string
	PageProperty     float64
}

func (m *MapFloat) Value() float64 {
	return m.PageProperty
}

func (m *MapFloat) UnmarshalJSON(data []byte) error {
	var err error
	var v interface{}
	if err = json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch x := v.(type) {
	case float64:
		m.PageProperty = x
	case map[string]interface{}:
		err = json.Unmarshal(data, &m.DatabaseProperty)
	case nil:
		m.IsNull = true
	default:
		err = fmt.Errorf("json: cannot unmarshal %v into Go value of type MapFloat", reflect.TypeOf(v).Name())
	}
	return err
}

func (m MapFloat) MarshalJSON() ([]byte, error) {
	if m.IsNull {
		return []byte("null"), nil
	}
	return json.Marshal(m.Value())
}
