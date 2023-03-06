package item

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Number struct {
	Value  float64
	Format struct {
		Format string `json:"format"`
	}
}

type NullNumber struct {
	Number Number
	Valid  bool
}

func (t NullNumber) ValueOrZero() Number {
	if !t.Valid {
		return Number{}
	}
	return t.Number
}

func (t *NullNumber) UnmarshalJSON(data []byte) error {
	var err error
	var v interface{}
	if err = json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch x := v.(type) {
	case Number:
		t.Number = x
	case map[string]interface{}:
		err = json.Unmarshal(data, &t.Number.Format)
	case interface{}:
		err = json.Unmarshal(data, &t.Number.Value)
	case nil:
		t.Valid = false
		return nil
	default:
		err = fmt.Errorf("json: cannot unmarshal %v into Go value of type NullNumber", reflect.TypeOf(v).Name())
	}
	t.Valid = err == nil
	return err
}

func (t NullNumber) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(t.Number.Value)
}
