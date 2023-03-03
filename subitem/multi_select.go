package subitem

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type MultiSelect struct {
	Selected []Select
	Options  []Select `json:"options"`
}

type NullMultiSelect struct {
	MultiSelect MultiSelect
	Valid       bool
}

func (s MultiSelect) GetAllPlainText() string {
	plainText := ""
	for i, value := range s.Selected {
		plainText += value.Name
		if i != len(s.Selected)-1 {
			plainText += ", "
		}
	}
	return plainText
}

func (s MultiSelect) GetAllOptionName() []string {
	var optionNames []string
	for _, value := range s.Selected {
		optionNames = append(optionNames, value.Name)
	}
	return optionNames
}

func (s NullMultiSelect) ValueOrZero() MultiSelect {
	if !s.Valid {
		return MultiSelect{}
	}
	return s.MultiSelect
}

func (s *NullMultiSelect) UnmarshalJSON(data []byte) error {
	var err error
	var v interface{}
	if err = json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch x := v.(type) {
	case MultiSelect:
		s.MultiSelect = x
	case []interface{}:
		err = json.Unmarshal(data, &s.MultiSelect.Selected)
	case map[string]interface{}:
		err = json.Unmarshal(data, &s.MultiSelect)
	case nil:
		s.Valid = false
		return nil
	default:
		err = fmt.Errorf("json: cannot unmarshal %v into Go value of type NullMultiSelect", reflect.TypeOf(v).Name())
	}
	s.Valid = err == nil
	return err
}

func (s NullMultiSelect) MarshalJSON() ([]byte, error) {
	if !s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.ValueOrZero().GetAllOptionName())
}
