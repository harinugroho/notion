package properties

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Option struct {
	Id        string   `json:"id"`
	Name      string   `json:"name"`
	Color     string   `json:"color"`
	OptionIds []string `json:"option_ids"`
}

type Options struct {
	Options []Option `json:"options"`
	Groups  []Option `json:"groups"`
}

type MapSelect struct {
	IsNull           bool
	DatabaseProperty Options
	PageProperty     Option
}

func (m *MapSelect) Value() string {
	return m.PageProperty.Name
}

func (m *MapSelect) UnmarshalJSON(data []byte) error {
	var err error
	var v interface{}
	if err = json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch x := v.(type) {
	case Option:
		m.PageProperty = x
	case map[string][]interface{}:
		err = json.Unmarshal(data, &m.DatabaseProperty)
	case map[string]interface{}:
		err = json.Unmarshal(data, &m.PageProperty)
	case nil:
		m.IsNull = true
	default:
		err = fmt.Errorf("json: cannot unmarshal %v into Go value of type MapSelect", reflect.TypeOf(v).Name())
	}
	return err
}

func (m MapSelect) MarshalJSON() ([]byte, error) {
	if m.IsNull {
		return []byte("null"), nil
	}
	return json.Marshal(m.Value())
}
