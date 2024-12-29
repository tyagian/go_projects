package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Example API URL
	url := "https://jsonplaceholder.typicode.com/posts/1"

	// Make the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		return
	}

	// Unmarshal into a map (dynamic JSON)
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	// Access specific fields
	fmt.Println("Title:", result["title"])
	fmt.Println("Body:", result["body"])
}
