package dao

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEmployeeError(t *testing.T) {
	emp, err := GetEmployee(123)
	assert.Nil(t, emp)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Status)
	assert.EqualValues(t, "Employee 123 was not found", err.Message)
	// if emp != nil {
	// 	t.Error("emp_id with 123 is not accepted")
	// }
	// if err == nil {
	// 	t.Error("excepting error for emp_id 123")
	// }
	// if err.StatusCode != http.StatusNotFound {
	// 	t.Error("Excepting 404 status when emp_id is 123")
	// }
}

func TestGetEmployeeNoError(t *testing.T) {
	emp, err := GetEmployee(101)
	assert.Nil(t, err)
	assert.NotNil(t, emp)
	assert.EqualValues(t, 200, http.StatusOK)
	assert.EqualValues(t, 101, emp.EmpId)
	assert.EqualValues(t, "Deepak", emp.EmpName)
	assert.EqualValues(t, "Deepak420@gmail.com", emp.EmpEmail)
	// if emp == nil {
	// 	t.Error("emp_id with 100 is accepted")
	// }
	// if err != nil {
	// 	t.Error("Not accepted input for emp_id")
	// }
}

/*
Output:
$ go test
PASS
ok      GO_Mini_Projects/golang-microservice/src/webapp/dao     0.218s

$ go test -cover
PASS
coverage: 100.0% of statements
ok      GO_Mini_Projects/golang-microservice/src/webapp/dao     1.361s
*/
