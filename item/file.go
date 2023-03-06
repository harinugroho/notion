package item

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

type Files []struct {
	Name string `json:"name"`
	Type string `json:"type"`
	File struct {
		Url        string    `json:"url"`
		ExpiryTime time.Time `json:"expiry_time"`
	} `json:"file"`
}

type NullFiles struct {
	Files Files
	Valid bool
}

func (s Files) GetAllPlainText() string {
	plainText := ""
	for i, value := range s {
		plainText += value.Name
		if i != len(s)-1 {
			plainText += ", "
		}
	}
	return plainText
}

func (s Files) GetAllUrl() []string {
	var urls []string
	for _, value := range s {
		urls = append(urls, value.File.Url)
	}
	return urls
}

func (t NullFiles) ValueOrZero() Files {
	if !t.Valid {
		return Files{}
	}
	return t.Files
}

func (t *NullFiles) UnmarshalJSON(data []byte) error {
	var err error
	var v interface{}
	if err = json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch x := v.(type) {
	case Files:
		t.Files = x
	case []interface{}:
		err = json.Unmarshal(data, &t.Files)
	case map[string]interface{}:
	case nil:
		t.Valid = false
		return nil
	default:
		err = fmt.Errorf("json: cannot unmarshal %v into Go value of type NullFile", reflect.TypeOf(v).Name())
	}
	t.Valid = err == nil
	return err
}

func (t NullFiles) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(t.ValueOrZero().GetAllUrl())
}
