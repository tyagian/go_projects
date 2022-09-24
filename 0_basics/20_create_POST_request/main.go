package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb video")
	PerformPostJsonRequest()
}

func PerformGetRequest() {

	//
}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:1111/post"

	// fake json payload

	requestBody := strings.NewReader(`
		{
			"course": "Let's go with Golang",
			"price": 0,
			"platform": "github.com"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
