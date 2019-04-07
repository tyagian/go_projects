package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	resp, err := http.Post("https://httpbin.org/post", "text/plain",
		strings.NewReader("this is the requested content"))
	if err != nil {
		log.Fatalln("Unable to get response")
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("unable to read content")
	}
	fmt.Println(string(content))
}
