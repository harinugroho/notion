package properties

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type MapMultiSelect struct {
	DatabaseProperty Options
	PageProperty     []Option
}

func (m *MapMultiSelect) Value() []string {
	var values []string
	for _, v := range m.PageProperty {
		values = append(values, v.Name)
	}
	return values
}

func (m *MapMultiSelect) UnmarshalJSON(data []byte) error {
	var err error
	var v interface{}
	if err = json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch x := v.(type) {
	case []Option:
		m.PageProperty = x
	case map[string]interface{}:
		err = json.Unmarshal(data, &m.DatabaseProperty)
	case []interface{}:
		err = json.Unmarshal(data, &m.PageProperty)
	default:
		err = fmt.Errorf("json: cannot unmarshal %v into Go value of type MapMultiSelect", reflect.TypeOf(v).Name())
	}
	return err
}

func (m MapMultiSelect) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.Value())
}
