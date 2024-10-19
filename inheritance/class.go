package main

import "fmt"

type Person struct {
	name string
	edad int
}

type EmployeeV3 struct {
	id int
}

type FullTimeEmployee struct {
	Person
	EmployeeV3
	endDate string
}

type TemporaryEmployee struct {
	Person
	EmployeeV3
	taxRate int
}

func (ftEmployee FullTimeEmployee) getMessage() string {
	return "Full time Employee"
}

func (ftEmployee TemporaryEmployee) getMessage() string {
	return "Temporary Employee"
}

func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

type PrintInfo interface {
	getMessage() string
}

func main() {
	ftEmployee := FullTimeEmployee{}

	ftEmployee.name = "ser"
	ftEmployee.edad = 12
	ftEmployee.id = 5

	tEmployee := TemporaryEmployee{}
	getMessage(tEmployee)
	getMessage(ftEmployee)
}
