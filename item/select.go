package item

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Select struct {
	Id        string   `json:"id"`
	Name      string   `json:"name"`
	Color     string   `json:"color"`
	Options   []Select `json:"options"`
	OptionIds []string `json:"option_ids"` // for group
}

type NullSelect struct {
	Select Select
	Valid  bool
}

func (s NullSelect) ValueOrZero() Select {
	if !s.Valid {
		return Select{}
	}
	return s.Select
}

func (s *NullSelect) UnmarshalJSON(data []byte) error {
	var err error
	var v interface{}
	if err = json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch x := v.(type) {
	case Select:
		s.Select = x
	case map[string]interface{}:
		err = json.Unmarshal(data, &s.Select)
	case nil:
		s.Valid = false
		return nil
	default:
		err = fmt.Errorf("json: cannot unmarshal %v into Go value of type NullSelect", reflect.TypeOf(v).Name())
	}
	s.Valid = err == nil
	return err
}

func (s NullSelect) MarshalJSON() ([]byte, error) {
	if !s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.Select.Name)
}
