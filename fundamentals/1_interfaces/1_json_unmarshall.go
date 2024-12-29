package main

// JSON Unmarshalling into a Struct with Embedded Interface
// Sometimes you may want to unmarshal JSON into a struct but allow
// flexibility for certain fields.
import (
	"encoding/json"
	"fmt"
)

type Data struct {
	Name string      `json:"name"`
	Info interface{} `json:"info"` // Flexible field
}

func main() {
	jsonStr := `{"name": "example", "info": {"age": 25, "city": "New York"}}`

	var data Data
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Name:", data.Name)
	fmt.Println("Info:", data.Info)

	// Type assertion to access Info fields
	if infoMap, ok := data.Info.(map[string]interface{}); ok {
		fmt.Println("Age:", infoMap["age"])
		fmt.Println("City:", infoMap["city"])
	}
}
