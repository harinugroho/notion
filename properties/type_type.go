package properties

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type MapType struct {
	DatabaseProperty map[string]string
	PageProperty     []Type
}

func (m *MapType) Value(separator string) string {
	return joinValueType(m.PageProperty, separator)
}

func joinValueType(types []Type, separator string) string {
	var value string
	for i, t := range types {
		value += fmt.Sprintf("%v", t.Value())
		if i != len(types)-1 {
			value += separator
		}
	}
	return value
}

func (m *MapType) ArrayValue() []string {
	return joinArrayValueType(m.PageProperty)
}

func joinArrayValueType(types []Type) []string {
	var values []string
	for _, v := range types {
		values = append(values, fmt.Sprintf("%v", v.Value()))
	}
	return values
}

func (m *MapType) UnmarshalJSON(data []byte) error {
	var err error
	var v interface{}
	if err = json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch x := v.(type) {
	case []Type:
		m.PageProperty = x
	case []interface{}:
		err = json.Unmarshal(data, &m.PageProperty)
	case map[string]interface{}:
		err = json.Unmarshal(data, &m.DatabaseProperty)
	default:
		err = fmt.Errorf("json: cannot unmarshal %v into Go value of type MapType", reflect.TypeOf(v).Name())
	}
	return err
}

func (m MapType) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.Value(" "))
}
