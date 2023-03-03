package subitem

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type NullString struct {
	String string
	Valid  bool
}

func (s NullString) ValueOrZero() string {
	if !s.Valid {
		return ""
	}
	return s.String
}

func (s *NullString) UnmarshalJSON(data []byte) error {
	var err error
	var v interface{}
	if err = json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch x := v.(type) {
	case string:
		s.String = x
	case map[string]interface{}:
	case nil:
		s.Valid = false
		return nil
	default:
		err = fmt.Errorf("json: cannot unmarshal %v into Go value of type NullString", reflect.TypeOf(v).Name())
	}
	s.Valid = err == nil
	return err
}

func (s NullString) MarshalJSON() ([]byte, error) {
	if !s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.String)
}
