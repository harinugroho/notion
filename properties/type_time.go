package properties

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

type MapTime struct {
	IsNull           bool
	DatabaseProperty map[string]string
	PageProperty     time.Time
}

func (m *MapTime) Value() time.Time {
	return m.PageProperty
}

func (m *MapTime) UnmarshalJSON(data []byte) error {
	var err error
	var v interface{}
	if err = json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch x := v.(type) {
	case time.Time:
		m.PageProperty = x
	case string:
		m.PageProperty, err = time.Parse(time.RFC3339, x)
	case map[string]interface{}:
		err = json.Unmarshal(data, &m.DatabaseProperty)
	case nil:
		m.IsNull = true
	default:
		err = fmt.Errorf("json: cannot unmarshal %v into Go value of type MapTime", reflect.TypeOf(v).Name())
	}
	return err
}

func (m MapTime) MarshalJSON() ([]byte, error) {
	if m.IsNull {
		return []byte("null"), nil
	}
	return json.Marshal(m.Value())
}
