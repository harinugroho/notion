package item

import (
	"encoding/json"
	"notion/subitem"
	"strconv"
)

type Type struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`

	Checkbox       subitem.NullCheckbox    `json:"checkbox"`
	PhoneNumber    subitem.NullString      `json:"phone_number"`
	Email          subitem.NullString      `json:"email"`
	Url            subitem.NullString      `json:"url"`
	PageId         subitem.NullString      `json:"page_id"`
	CreatedTime    subitem.NullString      `json:"created_time"`
	LastEditedTime subitem.NullString      `json:"last_edited_time"`
	Select         subitem.NullSelect      `json:"select"`
	MultiSelect    subitem.NullMultiSelect `json:"multi_select"`
	RichText       subitem.NullRichText    `json:"rich_text"`
	Title          subitem.NullRichText    `json:"title"`
	Number         subitem.NullNumber      `json:"number"`
	Status         subitem.NullStatus      `json:"status"`
	Date           subitem.NullDate        `json:"date"`
	People         subitem.NullPeople      `json:"people"`
	Files          subitem.NullFiles       `json:"files"`
	Formula        *Type                   `json:"formula"`
}

func (t Type) GetPlainData() string {
	switch t.Type {
	case "title":
		return t.Title.ValueOrZero()
	case "rich_text":
		return t.RichText.ValueOrZero()
	case "select":
		return t.Select.ValueOrZero().Name
	case "multi_select":
		return t.MultiSelect.ValueOrZero().GetAllPlainText()
	case "status":
		return t.Status.ValueOrZero().Selected.Name
	case "number":
		return strconv.FormatFloat(t.Number.ValueOrZero().Value, 'f', -1, 64)
	case "date":
		return t.Date.ValueOrZero()
	case "people":
		return t.People.ValueOrZero().GetAllPlainText()
	case "checkbox":
		//boolean to string
		return strconv.FormatBool(t.Checkbox.Checkbox)
	case "files":
		return t.Files.ValueOrZero().GetAllPlainText()
	case "formula":
		return t.Formula.GetPlainData()
	case "created_time":
		return t.CreatedTime.String
	case "url":
		return t.Url.String
	case "email":
		return t.Email.String
	case "phone_number":
		return t.PhoneNumber.String
	case "page":
		return t.PageId.String
	default:
		return ""
	}
}

func (t Type) MarshalJSON() ([]byte, error) {
	switch t.Type {
	case "title":
		return json.Marshal(t.Title)
	case "rich_text":
		return json.Marshal(t.RichText)
	case "select":
		return json.Marshal(t.Select)
	case "multi_select":
		return json.Marshal(t.MultiSelect)
	case "status":
		return json.Marshal(t.Status)
	case "number":
		return json.Marshal(t.Number)
	case "date":
		return json.Marshal(t.Date)
	case "people":
		return json.Marshal(t.People)
	case "checkbox":
		return json.Marshal(t.Checkbox)
	case "files":
		return json.Marshal(t.Files)
	case "formula":
		return json.Marshal(t.Formula)
	case "created_time":
		return json.Marshal(t.CreatedTime)
	case "url":
		return json.Marshal(t.Url)
	case "email":
		return json.Marshal(t.Email)
	case "phone_number":
		return json.Marshal(t.PhoneNumber)
	case "page":
		return json.Marshal(t.PageId)
	default:
		return []byte("null"), nil
	}
}
