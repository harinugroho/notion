package notion

import (
	"fmt"
	"testing"
)

func TestNewDatabase_Info(t *testing.T) {
	token := "secret_YoQ2x14O0UvVb0iXoK2DiAg3vIMA25oyquO70441n4w"
	id := "bc780c449ce74079b2a6c44425241aa4"
	fmt.Println("TOKEN: ", token)
	database := NewDatabase(token, id)
	info := database.GetInfo()
	for key, value := range info {
		fmt.Printf("%s: %s\n", key, value)
	}
}

func TestNewDatabase_Columns(t *testing.T) {
	token := "secret_YoQ2x14O0UvVb0iXoK2DiAg3vIMA25oyquO70441n4w"
	id := "bc780c449ce74079b2a6c44425241aa4"
	fmt.Println("TOKEN: ", token)
	database := NewDatabase(token, id)
	columns := database.GetColumns()
	for key, value := range columns {
		fmt.Printf("%s: %s\n", key, value)
	}
}
