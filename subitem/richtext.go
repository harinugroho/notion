package subitem

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type RichText []struct {
	Type        string `json:"type"`
	Text        Text   `json:"text"`
	Annotations struct {
		Bold          bool   `json:"bold"`
		Italic        bool   `json:"italic"`
		Strikethrough bool   `json:"strikethrough"`
		Underline     bool   `json:"underline"`
		Code          bool   `json:"code"`
		Color         string `json:"color"`
	} `json:"annotations"`
	PlainText string `json:"plain_text"`
	Href      string `json:"href"`
}

type NullRichText struct {
	RichText RichText
	Valid    bool
}

func (r RichText) getAllPlainText() string {
	plainText := ""
	for i, value := range r {
		plainText += value.PlainText
		if i != len(r)-1 {
			plainText += " "
		}
	}
	return plainText
}

func (r NullRichText) ValueOrZero() string {
	if !r.Valid || len(r.RichText) == 0 {
		return ""
	}
	return r.RichText.getAllPlainText()
}

func (r *NullRichText) UnmarshalJSON(data []byte) error {
	var err error
	var v interface{}
	if err = json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch x := v.(type) {
	case RichText:
		r.RichText = x
	case []interface{}:
		err = json.Unmarshal(data, &r.RichText)
	case map[string]interface{}:
	case nil:
		r.Valid = false
		return nil
	default:
		fmt.Println(x)
		err = fmt.Errorf("json: cannot unmarshal %v into Go value of type NullRichText", reflect.TypeOf(v).Name())
	}
	r.Valid = err == nil
	return err
}

func (r NullRichText) MarshalJSON() ([]byte, error) {
	if !r.Valid || len(r.RichText) == 0 {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf(`"%s"`, r.RichText.getAllPlainText())), nil
}
