package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJson(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestEncodeJSON(t *testing.T) {
	logJson("taufiq")
	logJson(1)
	logJson(true)
	logJson([]string{"taufiq", "kurniawan"})
}
