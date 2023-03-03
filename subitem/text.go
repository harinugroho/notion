package subitem

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Text struct {
	Content string      `json:"content"`
	Link    interface{} `json:"link"`
}

type NullText struct {
	Text  Text
	Valid bool
}

func (t NullText) ValueOrZero() Text {
	if !t.Valid {
		return Text{}
	}
	return t.Text
}

func (t *NullText) UnmarshalJSON(data []byte) error {
	var err error
	var v interface{}
	if err = json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch x := v.(type) {
	case Text:
		t.Text = x
	case map[string]interface{}:
		err = json.Unmarshal(data, &t.Text)
	case nil:
		t.Valid = false
		return nil
	default:
		err = fmt.Errorf("json: cannot unmarshal %v into Go value of type NullText", reflect.TypeOf(v).Name())
	}
	t.Valid = err == nil
	return err
}

func (t NullText) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(t.Text)
}
