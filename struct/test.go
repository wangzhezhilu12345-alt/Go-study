package main

import "time"

func main() {
	type Employee struct {
		ID      int
		Name    string
		Address string
		Dob     time.Time
		Salary  int
	}

	_ = Employee{}
}
