package item

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type People []*User

type NullPeople struct {
	People People
	Valid  bool
}

func (s People) GetAllPlainText() string {
	plainText := ""
	for i, person := range s {
		plainText += person.Name
		if i != len(s)-1 {
			plainText += ", "
		}
	}
	return plainText
}

func (s People) GetAllId() []string {
	var names []string
	for _, value := range s {
		names = append(names, value.Id)
	}
	return names
}

func (t NullPeople) ValueOrZero() People {
	if !t.Valid {
		return People{}
	}
	return t.People
}

func (t *NullPeople) UnmarshalJSON(data []byte) error {
	var err error
	var v interface{}
	if err = json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch x := v.(type) {
	case People:
		t.People = x
	case []interface{}:
		err = json.Unmarshal(data, &t.People)
	case map[string]interface{}:
	case nil:
		t.Valid = false
		return nil
	default:
		err = fmt.Errorf("json: cannot unmarshal %v into Go value of type NullPeople", reflect.TypeOf(v).Name())
	}
	t.Valid = err == nil
	return err
}

func (t NullPeople) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(t.People.GetAllId())
}
