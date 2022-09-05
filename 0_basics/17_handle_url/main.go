package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghnj456"

func main() {

	fmt.Println("Welcome to handling urls in Golang")
	fmt.Println(myurl)

	// parse url
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	// Port is a method not property
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	// url.Values
	fmt.Printf("The type of query params are: %T \n", qparams)
	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("Params is:", val)
	}

	// construct a web request
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}
