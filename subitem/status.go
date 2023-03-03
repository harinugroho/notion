package subitem

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Status struct {
	Selected Select
	Options  []Select `json:"options"`
	Groups   []Select `json:"groups"`
}

type NullStatus struct {
	Status Status
	Valid  bool
}

func (s NullStatus) ValueOrZero() Status {
	if !s.Valid {
		return Status{}
	}
	return s.Status
}

func (s *NullStatus) UnmarshalJSON(data []byte) error {
	var err error
	var v interface{}
	if err = json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch x := v.(type) {
	case Status:
		s.Status = x
	case map[string][]interface{}:
		err = json.Unmarshal(data, &s.Status)
	case interface{}:
		err = json.Unmarshal(data, &s.Status.Selected)
	case nil:
		s.Valid = false
		return nil
	default:
		err = fmt.Errorf("json: cannot unmarshal %v into Go value of type NullStatus", reflect.TypeOf(v).Name())
	}
	s.Valid = err == nil
	return err
}

func (s NullStatus) MarshalJSON() ([]byte, error) {
	if !s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.Status.Selected.Name)
}
