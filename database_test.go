package notion

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestClient_Database(t *testing.T) {
	client, err := NewClient(
		"secret_YoQ2x14O0UvVb0iXoK2DiAg3vIMA25oyquO70441n4w",
	).SetDatabaseIdByUrl(
		"https://www.notion.so/indieapps/bc780c449ce74079b2a6c44425241aa4?v=c13d78b65c8a4cde8d25c660bae94ebf&pvs=4",
	).GetDatabase()

	if err != nil {
		panic(err)
	}

	fmt.Println("============    Info    ============")
	info := client.GetInfo()
	for key, value := range info {
		fmt.Printf("%s:%v\n", key, value)
	}

	fmt.Println("============ Properties ============")
	listProperties := client.GetProperties()
	for key, value := range listProperties {
		fmt.Printf("%s:%v\n", key, value)
	}
}

func TestClient_DatabaseJson(t *testing.T) {
	client, err := NewClient(
		"secret_YoQ2x14O0UvVb0iXoK2DiAg3vIMA25oyquO70441n4w",
	).SetDatabaseIdByUrl(
		"https://www.notion.so/indieapps/bc780c449ce74079b2a6c44425241aa4?v=c13d78b65c8a4cde8d25c660bae94ebf&pvs=4",
	).GetDatabase()

	if err != nil {
		panic(err)
	}

	fmt.Println("============  Database  ============")
	info := client.GetObject()
	indent, err := json.MarshalIndent(info, "", "\t")
	if err != nil {
		return
	}
	fmt.Println(string(indent))
}

func TestList_Footer(t *testing.T) {
	client, err := NewClient(
		"secret_YoQ2x14O0UvVb0iXoK2DiAg3vIMA25oyquO70441n4w",
	).SetDatabaseIdByUrl(
		"https://www.notion.so/indieapps/bc780c449ce74079b2a6c44425241aa4?v=c13d78b65c8a4cde8d25c660bae94ebf&pvs=4",
	).GetList()
	if err != nil {
		panic(err)
	}

	fmt.Println("============   Footer   ============")
	results := client.GetFooter(map[string]string{
		"Kolom Number": "sum",
	})
	for key, result := range results {
		fmt.Printf("%s:%v\n", key, result)
	}
	fmt.Println("====================================")
}
