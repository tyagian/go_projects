package main

import (
	"encoding/json"
	"fmt"
)

// not exporting course so even with lower case is fine other use
// as Course
type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Printf("example of JSON creation")
	EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	CoursesList := []course{
		{"ReactJS Bootcamp", 299, "AWS", "abc123", []string{"webdev", "js"}},
		{"MERN Bootcamp", 199, "AWS", "bcd123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 299, "AWS", "hit123", nil},
	}

	// package this data as JSON data
	finalJson, err := json.MarshalIndent(CoursesList, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)

}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJS Bootcamp",
		"Price": 299,
		"website": "LearnCodeOnline.in",
		"tags": ["web-dev","js"]
	}
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		// %#v to print interface
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	// some cases where you just want to add data to key value

	// why use interface?
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v  and value is %v and Type is: %T\n", k, v, v)
	}

}
