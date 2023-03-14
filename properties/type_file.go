package properties

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

type File struct {
	Url        string    `json:"url"`
	ExpiryTime time.Time `json:"expiry_time"`
}

type MapFile struct {
	DatabaseProperty map[string]string
	PageProperty     File
}

func (m *MapFile) Value() string {
	return m.PageProperty.Url
}

func (m *MapFile) UnmarshalJSON(data []byte) error {
	var err error
	var v interface{}
	if err = json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch x := v.(type) {
	case File:
		m.PageProperty = x
	case map[string]interface{}:
		err = json.Unmarshal(data, &m.PageProperty)
	default:
		err = fmt.Errorf("json: cannot unmarshal %v into Go value of type MapFile", reflect.TypeOf(v).Name())
	}
	return err
}

func (m MapFile) MarshalJSON() ([]byte, error) {
	if m.Value() == "" {
		return []byte("null"), nil
	}
	return json.Marshal(m.Value())
}
