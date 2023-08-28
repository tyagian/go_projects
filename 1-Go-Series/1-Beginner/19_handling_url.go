package main

// using net library
import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid-ghadd"

func main() {
	fmt.Println("handling urls example")
	fmt.Println(myurl)
	// parsing url
	result, _ := url.Parse(myurl)
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Println("The type of query params are: %T\n", qparams)
	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("params are: ", val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}
	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}
