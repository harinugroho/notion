package notion

import (
	"fmt"
	"testing"
)

func TestNewList_Data(t *testing.T) {
	token := "secret_YoQ2x14O0UvVb0iXoK2DiAg3vIMA25oyquO70441n4w"
	id := "bc780c449ce74079b2a6c44425241aa4"
	list := NewList(token, id)
	data := list.GetData()
	for index, datum := range data {
		fmt.Printf("=====Data %d=====\n", index)
		for key, value := range datum {
			fmt.Printf("%s: %s\n", key, value.GetPlainData())
		}
	}
}

func TestNewList_Summary(t *testing.T) {
	token := "secret_YoQ2x14O0UvVb0iXoK2DiAg3vIMA25oyquO70441n4w"
	id := "bc780c449ce74079b2a6c44425241aa4"
	list := NewList(token, id)

	request := map[string]string{
		"Kolom Angka": "sum",
	}
	summary := list.GetSummary(request)
	for key, value := range summary {
		fmt.Printf("%s: %f\n", key, value)
	}

	request = map[string]string{
		"Kolom Angka": "count",
	}
	summary = list.GetSummary(request)
	for key, value := range summary {
		fmt.Printf("%s: %f\n", key, value)
	}

	request = map[string]string{
		"Kolom Angka": "max",
	}
	summary = list.GetSummary(request)
	for key, value := range summary {
		fmt.Printf("%s: %f\n", key, value)
	}

	request = map[string]string{
		"Kolom Angka": "min",
	}
	summary = list.GetSummary(request)
	for key, value := range summary {
		fmt.Printf("%s: %f\n", key, value)
	}
}
