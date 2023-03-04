package notion

import (
	"encoding/json"
	"fmt"
	"github.com/harinugroho/notion/item"
	"io"
	"net/http"
)

type List struct {
	object item.Object
}

func (list *List) queryDatabase(token string, id string) error {
	url := fmt.Sprintf("https://api.notion.com/v1/databases/%s/query", id)
	bearer := fmt.Sprintf("Bearer %s", token)
	req, err := http.NewRequest("POST", url, nil)
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Notion-Version", "2022-02-22")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("error closing body: %s", err)
		}
	}(resp.Body)

	err = json.NewDecoder(resp.Body).Decode(&list.object)
	if err != nil {
		return err
	}

	return nil
}

func (list *List) GetData() []map[string]item.Type {
	data, err := list.object.GetList()
	if err != nil {
		return nil
	}
	return data
}

func (list *List) GetSummary(request map[string]string) map[string]float64 {
	data, err := list.object.GetList()
	if err != nil {
		return nil
	}

	summary := map[string]float64{}
	for key := range request {
		summary[key] = 0
	}

	for index, value := range data {
		for key, action := range request {
			if action == "sum" || action == "average" {
				summary[key] += value[key].GetNumberData()
			} else if action == "max" {
				if index == 0 {
					summary[key] = value[key].GetNumberData()
				}
				if value[key].GetNumberData() > summary[key] {
					summary[key] = value[key].GetNumberData()
				}
			} else if action == "min" {
				if index == 0 {
					summary[key] = value[key].GetNumberData()
				}
				if value[key].GetNumberData() < summary[key] {
					summary[key] = value[key].GetNumberData()
				}
			}
		}
	}

	for key, value := range request {
		if value == "average" {
			summary[key] = summary[key] / float64(len(data))
		} else if value == "count" {
			summary[key] = float64(len(data))
		}
	}

	return summary
}

func NewList(token string, id string) List {
	var list List
	err := list.queryDatabase(token, id)
	if err != nil {
		panic(err)
	}
	return list
}
