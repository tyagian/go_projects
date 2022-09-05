package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// https://golang.org/pkg/net/http/
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Homepage")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))

}
