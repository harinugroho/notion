package notion

import (
	"github.com/harinugroho/notion/properties"
	"regexp"
)

type Client struct {
	integrationToken string
	databaseId       string

	sorts   []map[string]string
	filters map[string][]map[string]interface{}
	object  properties.Object
}

func (c *Client) SetDatabaseId(databaseId string) *Client {
	c.databaseId = databaseId
	return c
}

func (c *Client) SetDatabaseIdByUrl(databaseUrl string) *Client {
	re := regexp.MustCompile(`/[A-Za-z0-9]+\?v`)
	match := re.FindStringSubmatch(databaseUrl)[0]
	databaseId := match[1 : len(match)-2]
	c.databaseId = databaseId
	return c
}

func (c *Client) GetObject() properties.Object {
	return c.object
}

func NewClient(integrationToken string) *Client {
	return &Client{
		integrationToken: integrationToken,
	}
}
