package errors

//struct for error message in JSON
type AppError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statuscode"`
	Status     string `json:"status"`
}
