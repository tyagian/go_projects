package model

type Employee struct {
	// converting struct to lower case
	EmpId    int64  `json:"empid"`
	EmpName  string `json:"emp_name"`
	EmpEmail string `json:"emp_email"`
}
