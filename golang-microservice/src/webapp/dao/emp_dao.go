package dao

import (
	"GO_Mini_Projects/golang-microservice/src/webapp/errors"
	"GO_Mini_Projects/golang-microservice/src/webapp/model"
	"fmt"
	"net/http"
)

var (
	employees = map[int64]*model.Employee{
		101: {EmpId: 101, EmpName: "Deepak", EmpEmail: "Deepak420@gmail.com"},
		102: {EmpId: 102, EmpName: "Anuj", EmpEmail: "Anuj@gmail.com"},
		103: {EmpId: 102, EmpName: "Swati", EmpEmail: "Swati@gmail.com"},
		104: {EmpId: 103, EmpName: "Saranya", EmpEmail: "Saranya@gmail.com"},
		105: {EmpId: 104, EmpName: "Reyansh", EmpEmail: "Reyansh@gmail.com"},
	}
)

func GetEmployee(empId int64) (*model.Employee, *errors.AppError) {
	if employee := employees[empId]; employee != nil {
		return employee, nil
	}
	//return nil, errors.New(fmt.Sprintf("Employee %v was not found", empId))
	return nil, &errors.AppError{
		Message:    fmt.Sprintf("Employee %v was not found", empId),
		StatusCode: http.StatusNotFound,
		Status:     "not_found",
	}
}
