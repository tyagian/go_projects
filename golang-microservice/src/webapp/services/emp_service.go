package services

import (
	"GO_Mini_Projects/golang-microservice/src/webapp/dao"

	"GO_Mini_Projects/golang-microservice/src/webapp/model"
)

func GetEmployee(empId int64) (*model.Employee, errors) {
	//return dao.GetEmployee(empId)
	employee, err := dao.GetEmployee(empId)
	if err != nil {
		return nil, err
	}
	return employee, nil
}
