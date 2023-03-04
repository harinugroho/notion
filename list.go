package notion

import (
	"encoding/json"
	"fmt"
	"github.com/harinugroho/notion/item"
	"io"
	"net/http"
)

type List struct {
	token  string
	id     string
	object item.Object
}

func (list *List) QueryDatabase() ([]map[string]item.Type, error) {
	url := fmt.Sprintf("https://api.notion.com/v1/databases/%s/query", list.id)
	bearer := fmt.Sprintf("Bearer %s", list.token)
	req, err := http.NewRequest("POST", url, nil)
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Notion-Version", "2022-02-22")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("error closing body: %s", err)
		}
	}(resp.Body)

	err = json.NewDecoder(resp.Body).Decode(&list.object)
	if err != nil {
		return nil, err
	}

	return list.object.GetList()
}

func NewList(token string, id string) List {
	return List{
		token: token,
		id:    id,
	}
}
