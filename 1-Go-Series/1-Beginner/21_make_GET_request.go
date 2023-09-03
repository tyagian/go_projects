package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("example of GET request")
	PerformaGetRequest()
}

func PerformaGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Status Code: ", response.StatusCode)
	fmt.Println("Content length: ", response.ContentLength)

	// variable responseString from strings.Builder package
	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	// we are converting here to string since content is in Byte format

	// one approach to get output
	//fmt.Println(string(content))

	// second approach to print response from body in string format
	// responseString has a method of .Write()
	byteCount, _ := responseString.Write(content)
	fmt.Println("ByteCount is: ", byteCount)
	fmt.Println(responseString.String())
}
