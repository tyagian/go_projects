package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// go mod tidy, go mod verify, go list, go list all, go list -m all,
// go list -m -versions github.com/gorilla/mux
// go mod why
// go mod graph
// go mod edit -go 1.16
// go mod edit -module 1.20
// go mod vendor
// go run -mod=vendor main.go
// https://go.dev/ref/mod

func main() {
	fmt.Println("example of JSON creation")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("Hey! mod users")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> mod example </h1>"))
}
