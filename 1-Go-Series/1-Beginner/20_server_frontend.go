package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("example of server frontend")
	PerformaGetRequest()
}

func PerformaGetRequest() {
	const myurl = "http://localhost:8000/get"
	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status Code:", response.StatusCode)
	fmt.Println("Status Code:", response.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)

	byteCount, _ := responseString.Write(content)

	//fmt.Println(content)
	//fmt.Println(string(content))

	fmt.Println("ByteCount is: ", byteCount)
	fmt.Println(responseString.String())
}
