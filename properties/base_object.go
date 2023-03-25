package properties

import (
	"encoding/json"
	"time"
)

type Object struct {
	Object string `json:"object"`

	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`

	Id             string          `json:"id"`
	Cover          Type            `json:"cover"`
	Icon           Type            `json:"icon"`
	CreatedTime    time.Time       `json:"created_time"`
	CreatedBy      *Object         `json:"created_by"`
	LastEditedBy   *Object         `json:"last_edited_by"`
	LastEditedTime time.Time       `json:"last_edited_time"`
	Title          []Type          `json:"title"`
	Description    []Type          `json:"description"`
	IsInline       bool            `json:"is_inline"`
	Properties     map[string]Type `json:"properties"`
	Parent         Type            `json:"parent"`
	Url            string          `json:"url"`
	Archived       bool            `json:"archived"`

	Results    []Object    `json:"results"`
	NextCursor interface{} `json:"next_cursor"`
	HasMore    bool        `json:"has_more"`
	Type       string      `json:"type"`
	Page       struct {
	} `json:"page"`
}

func (o *Object) MapInfo() map[string]interface{} {
	return map[string]interface{}{
		"id":               o.Id,
		"cover":            o.Cover.Value(),
		"icon":             o.Icon.Value(),
		"created_time":     o.CreatedTime.String(),
		"created_by":       o.CreatedBy.Id,
		"last_edited_by":   o.LastEditedBy.Id,
		"last_edited_time": o.LastEditedTime.String(),
		"title":            joinValueType(o.Title, ""),
		"description":      joinValueType(o.Description, ""),
		"is_inline":        o.IsInline,
		"parent":           o.Parent.Value(),
		"url":              o.Url,
		"archived":         o.Archived,
	}
}

func (o *Object) MapProperties() map[string]string {
	properties := map[string]string{}
	if o.Object == "database" {
		for key, value := range o.Properties {
			properties[key] = value.Type
		}
	} else if o.Object == "list" {
		for key, value := range o.Results[0].Properties {
			properties[key] = value.Type
		}
	}
	return properties
}

func (o *Object) MapResults() []map[string]interface{} {
	var results []map[string]interface{}
	for _, result := range o.Results {
		properties := map[string]interface{}{}
		for key, value := range result.Properties {
			properties[key] = value.Value()
		}
		results = append(results, properties)
	}
	return results
}

func (o Object) MarshalJSON() ([]byte, error) {
	switch o.Object {
	case "database":
		return json.Marshal(map[string]interface{}{
			"info":       o.MapInfo(),
			"properties": o.MapProperties(),
		})
	case "list":
		return json.Marshal(o.Results)
	case "page":
		return json.Marshal(o.Properties)
	case "error":
		return json.Marshal(map[string]interface{}{
			"status":  o.Status,
			"code":    o.Code,
			"message": o.Message,
		})
	}
	return json.Marshal(o)
}
