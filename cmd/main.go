package main

import (
	"fmt"
	"time"
)

type Employee struct {
	Name string
	Number int
	Boss *Employee
	HireDate time.Time
}

func main() {
	c := map[string]*Employee

	c
	e := Employee{
		Name: "Ashutosh", Number: 1, HireDate: time.Now(),
	}

	b := &Employee{ "Tripti", 2, nil, time.Now() }

	e.Boss = b
	fmt.Printf("%T %+[1]v\n", e)
	fmt.Printf("%T %+[1]v\n", b)


}
