package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {
	customer := Customer{
		FirstName:  "Taufiq",
		MiddleName: "Kurniawan",
		LastName:   "Bayu",
		Age:        23,
		Married:    true,
		Hobbies:    []string{"Gaming", "Reading", "Coding"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Taufiq","MiddleName":"Kurniawan","LastName":"Bayu","Age":23,"Married":true,"Hobbies":["Gaming","Reading","Coding"]}`

	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)

	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Hobbies)
}

// Test JSON Array Complex
func TestJSONArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName: "Taufiq",
		Addresses: []Address{
			{
				Street:     "Semarang Street 290",
				Country:    "Indonesia",
				PostalCode: "9999",
			},
			{
				Street:     "Bandung West 209",
				Country:    "Americas",
				PostalCode: "90000",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

// JSON Array Complex Decode
func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Taufiq","MiddleName":"","LastName":"","Age":0,"Married":false,"Hobbies":null,"Addresses":[{"Street":"Semarang Street 290","Country":"Indonesia","PostalCode":"9999"},{"Street":"Bandung West 209","Country":"Americas","PostalCode":"90000"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Addresses)
}

//Only JSON Array

func TestOnlyJSONArray(t *testing.T) {
	jsonString := `[{"Street":"Semarang Street 290","Country":"Indonesia","PostalCode":"9999"},{"Street":"Bandung West 209","Country":"Americas","PostalCode":"90000"}]`
	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}
	err := json.Unmarshal(jsonBytes, addresses)
	if err != nil {
		panic(err)
	}
	fmt.Println(addresses)
}

// Only Array JSON
func TestOnlyJSONArrayComplex(t *testing.T) {
	addresses := []Address{
		{
			Street:     "Semarang Street 290",
			Country:    "Indonesia",
			PostalCode: "9999",
		},
		{
			Street:     "Bandung West 209",
			Country:    "Americas",
			PostalCode: "90000",
		},
	}

	bytes, _ := json.Marshal(addresses)
	fmt.Println(string(bytes))
}
