package controllers

import (
	"GO_Mini_Projects/golang-microservice/src/webapp/errors"
	"GO_Mini_Projects/golang-microservice/src/webapp/services"
	"encoding/json"
	"net/http"
	"strconv"
)

func GetEmployee(response http.ResponseWriter, request *http.Request) {
	empId, err := strconv.ParseInt(request.URL.Query().Get("emp_id"), 10, 64)

	if err != nil {
		apiErr := &errors.AppError{
			Message:    "emp_id must be a number",
			StatusCode: http.StatusBadRequest,
			Status:     "bad_request",
		}
		jsonValue, _ := json.Marshal(apiErr)
		response.WriteHeader(apiErr.StatusCode)
		response.Write(jsonValue)
		//response.WriteHeader(http.StatusNotFound)
		//response.Write([]byte("emp_id must be a number"))
		return
	}
	emp, apiErr := services.GetEmployee(empId)
	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		response.WriteHeader(apiErr.StatusCode)
		response.Write([]byte(jsonValue))
		return
	}
	jsonValue, _ := json.Marshal(emp)
	response.Write(jsonValue)
}
