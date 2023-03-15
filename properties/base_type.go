package properties

import (
	"encoding/json"
)

type Annotations struct {
	Bold          bool   `json:"bold"`
	Italic        bool   `json:"italic"`
	Strikethrough bool   `json:"strikethrough"`
	Underline     bool   `json:"underline"`
	Code          bool   `json:"code"`
	Color         string `json:"color"`
}

type Type struct {
	Id   string `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`

	Rollup         *Type          `json:"rollup"`
	Title          MapType        `json:"title"`
	Files          MapType        `json:"files"`
	RichText       MapType        `json:"rich_text"`
	Array          MapType        `json:"array"`
	People         MapObject      `json:"people"`
	Number         MapFloat       `json:"number"`
	Email          MapString      `json:"email"`
	PhoneNumber    MapString      `json:"phone_number"`
	Url            MapString      `json:"url"`
	Checkbox       MapBoolean     `json:"checkbox"`
	CreatedTime    MapTime        `json:"created_time"`
	LastEditedTime MapTime        `json:"last_edited_time"`
	Text           MapText        `json:"text"`
	File           MapFile        `json:"file"`
	Formula        MapFormula     `json:"formula"`
	External       MapExternal    `json:"external"`
	Date           MapDate        `json:"date"`
	Select         MapSelect      `json:"select"`
	Status         MapSelect      `json:"status"`
	MultiSelect    MapMultiSelect `json:"multi_select"`
	Relation       MapRelation    `json:"relation"`
	CreatedBy      *Object        `json:"created_by"`
	LastEditedBy   *Object        `json:"last_edited_by"`

	Annotations Annotations `json:"annotations"`
	PlainText   string      `json:"plain_text"`
	Href        string      `json:"href"`

	PageId string `json:"page_id"`
	Emoji  string `json:"emoji"`
}

func (t *Type) Value() interface{} {
	switch t.Type {
	case "emoji":
		return t.Emoji
	case "external":
		return t.External.Value()
	case "file":
		return t.File.Value()
	case "formula":
		return t.Formula.Value()
	case "text":
		return t.PlainText
	case "page_id":
		return t.PageId
	case "checkbox":
		return t.Checkbox.Value()
	case "title":
		return t.Title.Value("")
	case "files":
		return t.Files.ArrayValue()
	case "rich_text":
		return t.RichText.Value("")
	case "array":
		return t.Array.ArrayValue()
	case "people":
		return t.People.ArrayValue()
	case "email":
		return t.Email.Value()
	case "phone_number":
		return t.PhoneNumber.Value()
	case "url":
		return t.Url.Value()
	case "number":
		return t.Number.Value()
	case "created_time":
		return t.CreatedTime.Value()
	case "last_edited_time":
		return t.LastEditedTime.Value()
	case "created_by":
		return t.CreatedBy.Id
	case "last_edited_by":
		return t.LastEditedBy.Id
	case "date":
		return t.Date.Value()
	case "select":
		return t.Select.Value()
	case "status":
		return t.Status.Value()
	case "multi_select":
		return t.MultiSelect.Value()
	case "rollup":
		return t.Rollup.Value()
	case "relation":
		return t.Relation.Value()
	default:
		return nil
	}
}

func (t Type) MarshalJSON() ([]byte, error) {
	switch t.Type {
	case "emoji":
		return json.Marshal(t.Emoji)
	case "external":
		return json.Marshal(t.External)
	case "file":
		return json.Marshal(t.File)
	case "formula":
		return json.Marshal(t.Formula)
	case "text":
		return json.Marshal(t.PlainText)
	case "page_id":
		return json.Marshal(t.PageId)
	case "checkbox":
		return json.Marshal(t.Checkbox)
	case "title":
		return json.Marshal(t.Title.ValueOrNull(""))
	case "files":
		return json.Marshal(t.Files)
	case "rich_text":
		return json.Marshal(t.RichText.ValueOrNull(""))
	case "array":
		return json.Marshal(t.Array)
	case "people":
		return json.Marshal(t.People)
	case "email":
		return json.Marshal(t.Email)
	case "phone_number":
		return json.Marshal(t.PhoneNumber)
	case "url":
		return json.Marshal(t.Url)
	case "number":
		return json.Marshal(t.Number)
	case "created_time":
		return json.Marshal(t.CreatedTime)
	case "last_edited_time":
		return json.Marshal(t.LastEditedTime)
	case "created_by":
		return json.Marshal(t.CreatedBy.Id)
	case "last_edited_by":
		return json.Marshal(t.LastEditedBy.Id)
	case "date":
		return json.Marshal(t.Date)
	case "select":
		return json.Marshal(t.Select)
	case "status":
		return json.Marshal(t.Status)
	case "multi_select":
		return json.Marshal(t.MultiSelect)
	case "rollup":
		return json.Marshal(t.Rollup)
	case "relation":
		return json.Marshal(t.Relation)
	default:
		return []byte("null"), nil
	}
}
