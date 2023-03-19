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
