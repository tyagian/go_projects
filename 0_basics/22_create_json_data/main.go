// convert data as valid json or encode data to json

package main

import (
	"encoding/json"
	"fmt"
)

// json:"-" to hide password
// json:"tags,omitempty" to omit or remove nil/null with

type course struct {
	Name     string `json: "coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON")
	EncodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "learncodeonline.in", "abc123", []string{"webdev", "js"}},
		{"MERN Bootcamp", 199, "learncodeonline.in", "bcd123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 299, "learncodeonline.in", "hit123", nil},
	}

	// package this data as JSON
	// MarshalIndent takes 2 parameters:
	// \t used for intendation
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}
