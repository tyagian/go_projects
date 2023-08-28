package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://example.com"

func main() {
	fmt.Println("web request example")
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Printf("web response type %T\n", response)
	// web response type *http.Response

	defer response.Body.Close()
	// caller's responsibility to close the connection

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)
}
