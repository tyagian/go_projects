package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	fmt.Println("Welcome to web verb video")
	PerformGetRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:1111/get"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code:", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)

	//fmt.Println(string(content))

	/*
		// Another way to print response in string format
		byteCount, _ := responseString.Write(content)
		fmt.Println("ByteCount is: ", byteCount)
		fmt.Println(responseString.String())
	*/
}
