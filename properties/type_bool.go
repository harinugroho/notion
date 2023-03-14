package properties

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type MapBoolean struct {
	DatabaseProperty map[string]string
	PageProperty     bool
}

func (m *MapBoolean) Value() bool {
	return m.PageProperty
}

func (m *MapBoolean) UnmarshalJSON(data []byte) error {
	var err error
	var v interface{}
	if err = json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch x := v.(type) {
	case bool:
		m.PageProperty = x
	case map[string]interface{}:
		err = json.Unmarshal(data, &m.DatabaseProperty)
	default:
		err = fmt.Errorf("json: cannot unmarshal %v into Go value of type MapBoolean", reflect.TypeOf(v).Name())
	}
	return err
}

func (m MapBoolean) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.Value())
}
