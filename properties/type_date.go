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
	return m.PageProperty.Start.Format(time.RFC3339) + " - " + m.PageProperty.End.Format(time.RFC3339)
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
		if x["start"] != nil {
			m.PageProperty.Start, err = time.Parse(time.RFC3339, x["start"].(string))
			if err != nil {
				m.PageProperty.Start, err = time.Parse("2006-01-02", x["start"].(string))
				if err != nil {
					return err
				}
			}
		}
		if x["end"] != nil {
			m.PageProperty.End, err = time.Parse(time.RFC3339, x["end"].(string))
			if err != nil {
				m.PageProperty.End, err = time.Parse("2006-01-02", x["end"].(string))
				if err != nil {
					return err
				}
			}
		}
		m.PageProperty.TimeZone = x["time_zone"]
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
