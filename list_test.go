package notion

import (
	"encoding/json"
	"fmt"
	"github.com/harinugroho/notion/properties"
	"testing"
)

func TestList_List(t *testing.T) {
	client, err := NewClient(
		"secret_YoQ2x14O0UvVb0iXoK2DiAg3vIMA25oyquO70441n4w",
	).SetDatabaseIdByUrl(
		"https://www.notion.so/indieapps/bc780c449ce74079b2a6c44425241aa4?v=c13d78b65c8a4cde8d25c660bae94ebf&pvs=4",
	).GetList()
	if err != nil {
		panic(err)
	}

	fmt.Println("============    List    ============")
	results := client.GetResults()
	for _, result := range results {
		for key, value := range result {
			fmt.Printf("%s:%v\n", key, value)
		}
		fmt.Println("====================================")
	}
}

func TestList_ListFilter(t *testing.T) {
	client, err := NewClient(
		"secret_YoQ2x14O0UvVb0iXoK2DiAg3vIMA25oyquO70441n4w",
	).SetDatabaseIdByUrl(
		"https://www.notion.so/indieapps/bc780c449ce74079b2a6c44425241aa4?v=c13d78b65c8a4cde8d25c660bae94ebf&pvs=4",
	).Filters(properties.Filter{
		Logic: "and",
		Items: []properties.FilterItem{
			{
				Property: "Kolom Checkbox",
				Type:     "checkbox",
				Logic:    "equals",
				Value:    true,
			},
		},
	}).GetList()
	if err != nil {
		panic(err)
	}

	fmt.Println("============    List    ============")
	results := client.GetResults()
	for _, result := range results {
		for key, value := range result {
			fmt.Printf("%s:%v\n", key, value)
		}
		fmt.Println("====================================")
	}
}

func TestList_ListShort(t *testing.T) {
	client, err := NewClient(
		"secret_YoQ2x14O0UvVb0iXoK2DiAg3vIMA25oyquO70441n4w",
	).SetDatabaseIdByUrl(
		"https://www.notion.so/indieapps/bc780c449ce74079b2a6c44425241aa4?v=c13d78b65c8a4cde8d25c660bae94ebf&pvs=4",
	).Sorts(map[string]string{
		"Kolom Number": "descending",
	}).GetList()
	if err != nil {
		panic(err)
	}

	fmt.Println("============    List    ============")
	results := client.GetResults()
	for _, result := range results {
		for key, value := range result {
			fmt.Printf("%s:%v\n", key, value)
		}
		fmt.Println("====================================")
	}
}

func TestList_ListFilterShort(t *testing.T) {
	client, err := NewClient(
		"secret_YoQ2x14O0UvVb0iXoK2DiAg3vIMA25oyquO70441n4w",
	).SetDatabaseIdByUrl(
		"https://www.notion.so/indieapps/bc780c449ce74079b2a6c44425241aa4?v=c13d78b65c8a4cde8d25c660bae94ebf&pvs=4",
	).Filters(properties.Filter{
		Logic: "and",
		Items: []properties.FilterItem{
			{
				Property: "Kolom Checkbox",
				Type:     "checkbox",
				Logic:    "equals",
				Value:    true,
			},
		},
	}).Sorts(map[string]string{
		"Kolom Number": "descending",
	}).GetList()
	if err != nil {
		panic(err)
	}

	fmt.Println("============    List    ============")
	results := client.GetResults()
	for _, result := range results {
		for key, value := range result {
			fmt.Printf("%s:%v\n", key, value)
		}
		fmt.Println("====================================")
	}
}

func TestList_ListJson(t *testing.T) {
	client, err := NewClient(
		"secret_YoQ2x14O0UvVb0iXoK2DiAg3vIMA25oyquO70441n4w",
	).SetDatabaseIdByUrl(
		"https://www.notion.so/indieapps/bc780c449ce74079b2a6c44425241aa4?v=c13d78b65c8a4cde8d25c660bae94ebf&pvs=4",
	).GetList()
	if err != nil {
		panic(err)
	}

	fmt.Println("============    List    ============")
	results := client.GetObject()
	indent, err := json.MarshalIndent(results, "", "\t")
	if err != nil {
		return
	}
	fmt.Println(string(indent))
}
