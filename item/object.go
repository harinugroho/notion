package item

import (
	"errors"
	"time"
)

type Object struct {
	Object         string          `json:"object"`
	Id             string          `json:"id"`
	CreatedTime    time.Time       `json:"created_time"`
	LastEditedTime time.Time       `json:"last_edited_time"`
	Title          NullRichText    `json:"title"`
	Results        []Object        `json:"results"`
	Properties     map[string]Type `json:"properties"`
	Parent         Type            `json:"parent"`
	Url            string          `json:"url"`
	Archived       bool            `json:"archived"`
}

func (o *Object) GetInfo() map[string]interface{} {
	details := make(map[string]interface{})
	details["id"] = o.Id
	details["url"] = o.Url
	details["title"] = o.Title.ValueOrZero()
	details["created_time"] = o.CreatedTime.Format("2006-01-02 15:04:05")
	details["last_edited_time"] = o.LastEditedTime.Format("2006-01-02 15:04:05")
	details["archived"] = o.Archived
	return details
}

func (o *Object) GetColumns() (map[string]string, error) {
	if o.Object == "database" {
		columns := make(map[string]string)
		for key, value := range o.Properties {
			columns[key] = value.Type
		}
		return columns, nil
	} else {
		return nil, errors.New("not a database")
	}
}

func (o *Object) GetList() ([]map[string]Type, error) {
	if o.Object == "list" {
		var data []map[string]Type
		for _, value := range o.Results {
			properties := value.Properties
			datum := map[string]Type{}
			for key, value := range properties {
				datum[key] = value
			}
			data = append(data, datum)
		}
		return data, nil
	} else {
		return nil, errors.New("not a list")
	}
}

func (o *Object) GetFooter() ([]map[string]Type, error) {
	if o.Object == "list" {
		var data []map[string]Type
		for _, value := range o.Results {
			properties := value.Properties
			datum := map[string]Type{}
			for key, value := range properties {
				datum[key] = value
			}
			data = append(data, datum)
		}
		return data, nil
	} else {
		return nil, errors.New("not a list")
	}
}
