package properties

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Relation struct {
	Id string `json:"id"`
}

type MapRelation struct {
	DatabaseProperty map[string]string
	PageProperty     []Relation
}

func (m *MapRelation) Value() []string {
	var values []string
	for _, v := range m.PageProperty {
		values = append(values, v.Id)
	}
	return values
}

func (m *MapRelation) UnmarshalJSON(data []byte) error {
	var err error
	var v interface{}
	if err = json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch x := v.(type) {
	case []Relation:
		m.PageProperty = x
	case []interface{}:
		err = json.Unmarshal(data, &m.PageProperty)
	case map[string]interface{}:
		m.DatabaseProperty = make(map[string]string)
		for key, value := range x {
			m.DatabaseProperty[key] = value.(string)
		}
	default:
		fmt.Println(string(data))
		err = fmt.Errorf("json: cannot unmarshal %v into Go value of type MapRelation", reflect.TypeOf(v).Name())
	}
	return err
}

func (m MapRelation) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.Value())
}
