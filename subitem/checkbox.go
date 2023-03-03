package subitem

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type NullCheckbox struct {
	Checkbox bool `json:"checkbox"`
	Valid    bool
}

func (t NullCheckbox) ValueOrZero() bool {
	if !t.Valid {
		return false
	}
	return t.Checkbox
}

func (t *NullCheckbox) UnmarshalJSON(data []byte) error {
	var err error
	var v interface{}
	if err = json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch x := v.(type) {
	case bool:
		t.Checkbox = x
	case map[string]interface{}:
	case nil:
		t.Valid = false
		return nil
	default:
		err = fmt.Errorf("json: cannot unmarshal %v into Go value of type NullCheckbox", reflect.TypeOf(v).Name())
	}
	t.Valid = err == nil
	return err
}

func (t NullCheckbox) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("%v", t.Checkbox)), nil
}
