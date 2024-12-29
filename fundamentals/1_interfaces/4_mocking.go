package main

import "fmt"

// Logger interface
type Logger interface {
	Log(message string)
}

// ConsoleLogger struct
type ConsoleLogger struct{}

// Implement Logger for ConsoleLogger
func (c ConsoleLogger) Log(message string) {
	fmt.Println("Console Log:", message)
}

// Function using Logger
func DoSomething(logger Logger) {
	logger.Log("Starting task...")
	// Task logic here
	logger.Log("Task completed.")
}

func main() {
	cl := ConsoleLogger{}
	DoSomething(cl)
}
