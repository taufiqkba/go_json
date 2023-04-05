package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJSON(t *testing.T) {
	jsonRequest := `{"FirstName": "Taufiq", "MiddleName": "Kurniawan", "LastName": "Bayu", "Age": 23}`
	jsonBytes := []byte(jsonRequest)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.LastName)
}
