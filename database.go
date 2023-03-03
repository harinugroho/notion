package notion

import (
	"NotionRest/notion/item"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Database struct {
	object item.Object
}

func (database *Database) retrieveDatabase(token string, id string) error {
	url := fmt.Sprintf("https://api.notion.com/v1/databases/%s", id)
	bearer := fmt.Sprintf("Bearer %s", token)
	req, err := http.NewRequest("GET", url, nil)
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

	err = json.NewDecoder(resp.Body).Decode(&database.object)
	if err != nil {
		return err
	}
	return nil
}

func (database *Database) GetInfo() map[string]interface{} {
	info := database.object.GetInfo()
	return info
}

func (database *Database) GetColumns() map[string]string {
	columns, err := database.object.GetColumns()
	if err != nil {
		panic(err)
	}
	return columns
}

func NewDatabase(token string, id string) Database {
	var database Database
	err := database.retrieveDatabase(token, id)
	if err != nil {
		panic(err)
	}
	return database
}
