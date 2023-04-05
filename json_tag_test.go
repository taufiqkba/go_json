package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func TestJSONTag(t *testing.T) {
	product := &Product{
		Id:       "P001",
		Name:     "Apple MacBook Pro M1 Pro",
		ImageURL: "https://example.com/image.png",
	}

	byte, _ := json.Marshal(product)
	fmt.Println(string(byte))
}

func TestJSONTagDecode(t *testing.T) {
	jsonString := `{"id":"P001","name":"Apple MacBook Pro M1 Pro","image_url":"https://example.com/image.png"}`
	jsonBytes := []byte(jsonString)

	product := &Product{}
	json.Unmarshal(jsonBytes, product)
	fmt.Println(product)
}
