package main

import "fmt"

type EmployeeV2 struct {
	id       int
	name     string
	vacation bool
}

func NewEmployee(id int, name string, vacation bool) *EmployeeV2 {
	return &EmployeeV2{
		id:       id,
		name:     name,
		vacation: vacation,
	}
}

func main() {
	e := EmployeeV2{}

	fmt.Println("%v\n", e)

	// e2 := EmployeeV2{
	// 	id:       1,
	// 	name:     "Sergio",
	// 	vacation: true,
	// }

	// e3 := new(EmployeeV2)

	// 4

	e4 := NewEmployee(1, "Ser", false)

	fmt.Println(e4)
}
