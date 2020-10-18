package dao

import (
	"net/http"
	"testing"
)

func TestGetEmployeeError(t *testing.T) {
	emp, err := GetEmployee(123)

	if emp != nil {
		t.Error("emp_id with 123 is not accepted")
	}
	if err == nil {
		t.Error("excepting error for emp_id 123")
	}
	if err.StatusCode != http.StatusNotFound {
		t.Error("Excepting 404 status when emp_id is 123")
	}
}

func TestGetEmployeeNoError(t *testing.T) {
	emp, err := GetEmployee(101)

	if emp == nil {
		t.Error("emp_id with 100 is accepted")
	}
	if err != nil {
		t.Error("Not accepted input for emp_id")
	}
}
