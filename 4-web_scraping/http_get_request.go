//web scraping with go

package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// make HTTP request
	response, err := http.Get("https://www.devdungeon.com/")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// copy data from the response to standard output
	n, err := io.Copy(os.Stdout, response.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Number of bytes copied to STDOUT", n)
}
