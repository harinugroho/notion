package properties

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

type Date struct {
	Start    time.Time   `json:"start"`
	End      time.Time   `json:"end"`
	TimeZone interface{} `json:"time_zone"`
}

type MapDate struct {
	IsNull           bool
	DatabaseProperty map[string]string
	PageProperty     Date
}

func (m *MapDate) Value() string {
	return m.PageProperty.Start.String() + " - " + m.PageProperty.End.String()
}

func (m *MapDate) UnmarshalJSON(data []byte) error {
	var err error
	var v interface{}
	if err = json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch x := v.(type) {
	case Date:
		m.PageProperty = x
	case map[string]interface{}:
		err = json.Unmarshal(data, &m.PageProperty)
	case nil:
		m.IsNull = true
	default:
		err = fmt.Errorf("json: cannot unmarshal %v into Go value of type MapDate", reflect.TypeOf(v).Name())
	}
	return err
}

func (m MapDate) MarshalJSON() ([]byte, error) {
	if m.IsNull {
		return []byte("null"), nil
	}
	return json.Marshal(m.Value())
}
