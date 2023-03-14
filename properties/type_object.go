package properties

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type MapObject struct {
	DatabaseProperty map[string]string
	PageProperty     []Object
}

func (m *MapObject) Value(separator string) string {
	return joinValueObject(m.PageProperty, separator)
}

func joinValueObject(types []Object, separator string) string {
	var value string
	for i, t := range types {
		value += fmt.Sprintf("%v", t.Id)
		if i != len(types)-1 {
			value += separator
		}
	}
	return value
}

func (m *MapObject) UnmarshalJSON(data []byte) error {
	var err error
	var v interface{}
	if err = json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch x := v.(type) {
	case []Object:
		m.PageProperty = x
	case []interface{}:
		err = json.Unmarshal(data, &m.PageProperty)
	case map[string]interface{}:
		err = json.Unmarshal(data, &m.DatabaseProperty)
	default:
		err = fmt.Errorf("json: cannot unmarshal %v into Go value of type MapObject", reflect.TypeOf(v).Name())
	}
	return err
}

func (m MapObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.Value(" "))
}
