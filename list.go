package notion

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/harinugroho/notion/properties"
	"io"
	"net/http"
)

func convertSorts(sorts map[string]string) []map[string]string {
	var sortsConverted []map[string]string
	for key, value := range sorts {
		sortsConverted = append(sortsConverted, map[string]string{
			"property":  key,
			"direction": value,
		})
	}
	return sortsConverted
}

func (c *Client) Sorts(sorts map[string]string) *Client {
	c.sorts = convertSorts(sorts)
	return c
}

func convertFilters(filter properties.Filter) map[string][]map[string]interface{} {
	filtersConverted := make(map[string][]map[string]interface{})
	items := make([]map[string]interface{}, 0)
	for _, value := range filter.Items {
		item := map[string]interface{}{
			"property": value.Property,
			value.Type: map[string]interface{}{
				value.Logic: value.Value,
			},
		}
		items = append(items, item)
	}
	filtersConverted[filter.Logic] = items
	return filtersConverted
}

func (c *Client) Filters(filters properties.Filter) *Client {
	c.filters = convertFilters(filters)
	return c
}

func (c *Client) GetList() (Client, error) {
	if c.databaseId == "" {
		panic("object id is empty")
	}

	url := fmt.Sprintf("https://api.notion.com/v1/databases/%s/query", c.databaseId)
	bearer := fmt.Sprintf("Bearer %s", c.integrationToken)

	bodyMap := map[string]interface{}{}
	if len(c.sorts) > 0 {
		bodyMap["sorts"] = c.sorts
	}
	if len(c.filters) > 0 {
		bodyMap["filter"] = c.filters
	}
	bodyJson, _ := json.Marshal(bodyMap)
	bodyReader := bytes.NewReader(bodyJson)

	req, err := http.NewRequest("POST", url, bodyReader)
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Notion-Version", "2022-02-22")
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return *c, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("error closing body: %s", err)
		}
	}(resp.Body)

	err = json.NewDecoder(resp.Body).Decode(&c.object)
	if err != nil {
		return *c, err
	}
	return *c, nil
}

func (c *Client) GetResults() []map[string]interface{} {
	return c.object.MapResults()
}
