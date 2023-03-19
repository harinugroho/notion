package notion

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetDatabase() (Client, error) {
	if c.databaseId == "" {
		panic("object id is empty")
	}

	url := fmt.Sprintf("https://api.notion.com/v1/databases/%s", c.databaseId)
	bearer := fmt.Sprintf("Bearer %s", c.integrationToken)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Notion-Version", "2022-02-22")

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

func (c *Client) GetInfo() map[string]interface{} {
	return c.object.MapInfo()
}

func (c *Client) GetProperties() map[string]string {
	return c.object.MapProperties()
}

func (c *Client) GetFooter(footers map[string]string) map[string]float64 {
	count := make(map[string]float64)
	result := make(map[string]float64)
	for key := range footers {
		count[key] = 0
		result[key] = 0
	}

	for _, value := range c.object.MapResults() {
		for key, action := range footers {
			count[key] += 1
			if action == "sum" {
				result[key] += value[key].(float64)
			} else if action == "count" {
				result[key] += 1
			} else if action == "avg" {
				result[key] += value[key].(float64)
			}
		}
	}

	for key, action := range footers {
		if action == "avg" {
			result[key] = result[key] / count[key]
		}
	}
	return result
}
