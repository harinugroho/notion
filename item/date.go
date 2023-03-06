package item

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Date struct {
	Start    string `json:"start"`
	End      string `json:"end"`
	TimeZone string `json:"time_zone"`
}

type NullDate struct {
	Date  Date
	Valid bool
}

func (s Date) getPlainText() string {
	plainText := ""
	if s.Start != "" {
		plainText += s.Start
	}
	if s.End != "" {
		plainText += " - " + s.End
	}
	return plainText
}

func (t NullDate) ValueOrZero() string {
	if !t.Valid {
		return ""
	}
	return t.Date.getPlainText()
}

func (t *NullDate) UnmarshalJSON(data []byte) error {
	var err error
	var v interface{}
	if err = json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch x := v.(type) {
	case Date:
		t.Date = x
	case map[string]interface{}:
		err = json.Unmarshal(data, &t.Date)
	case nil:
		t.Valid = false
		return nil
	default:
		err = fmt.Errorf("json: cannot unmarshal %v into Go value of type NullDate", reflect.TypeOf(v).Name())
	}
	t.Valid = err == nil
	return err
}

func (t NullDate) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return []byte("null"), nil
	}
	// convert date to bytes
	return json.Marshal(t.ValueOrZero())
}
