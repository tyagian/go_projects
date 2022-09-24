package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {

	fmt.Println("Welcome to web verb video")
	PerformPostFormRequest()
}

/*
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
//}

/*
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
} */

func PerformPostFormRequest() {

	const myurl = "http://localhost:1111/postform"

	// formdata
	data := url.Values{}
	data.Add("firstname", "anuj")
	data.Add("firstname", "tyagi")
	data.Add("email", "anuj@go.dev")

	response, err := http.PostForm(myurl, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))

}
