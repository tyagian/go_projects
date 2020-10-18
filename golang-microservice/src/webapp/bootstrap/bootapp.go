package bootstrap

import (
	"GO_Mini_Projects/golang-microservice/src/webapp/controllers"
	"net/http"
)

func BootApplication() {

	//fmt.Println("Bootstrap application package")
	http.HandleFunc("/employees", controllers.GetEmployee)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}
