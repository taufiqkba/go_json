package go_json

import (
	"encoding/json"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("CustomerOut.json")
	encoder := json.NewEncoder(writer)

	customer := &Customer{
		FirstName:  "Taufiq",
		MiddleName: "Kurniawan",
		LastName:   "Bayu",
	}

	encoder.Encode(customer)
}
