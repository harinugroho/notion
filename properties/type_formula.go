package properties

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type MapFormula struct {
	DatabaseProperty map[string]string
	PageProperty     *Type
}

func (m *MapFormula) Value() interface{} {
	return m.PageProperty.Value()
}

func (m *MapFormula) UnmarshalJSON(data []byte) error {
	var err error
	var v interface{}
	if err = json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch x := v.(type) {
	case *Type:
		m.PageProperty = x
	case map[string]string:
		err = json.Unmarshal(data, &m.DatabaseProperty)
	case map[string]interface{}:
		err = json.Unmarshal(data, &m.PageProperty)
	default:
		err = fmt.Errorf("json: cannot unmarshal %v into Go value of type MapFormula", reflect.TypeOf(v).Name())
	}
	return err
}

func (m MapFormula) MarshalJSON() ([]byte, error) {
	if m.Value() == "" {
		return []byte("null"), nil
	}
	return json.Marshal(m.Value())
}
