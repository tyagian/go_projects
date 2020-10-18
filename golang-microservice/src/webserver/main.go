package main

import "net/http"

func main() {
	// Handle (param-1 /endpoint, part-2 func as argument)
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello go!"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}
