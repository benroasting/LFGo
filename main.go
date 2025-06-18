package main

import (
	"fmt"
)

func main() {
// Variables

	var name string  = "Ben"
	fmt.Printf("Hello, %s!\n", name)

	age := 38
	fmt.Printf("You are %d years old.\n", age)

	var city string
	city = "Columbia"
	fmt.Printf("You live in %s.\n", city)

	// Zero values
	var defaultInt int
	var defaultFloat float64
	var defaultString string
	var defaultBool bool
	fmt.Printf("Default int: %d, float: %f, string: '%s', bool: %t\n", defaultInt, defaultFloat, defaultString, defaultBool) 

	var (
		isEmployed bool = true
		salary int = 75000
		position string = "Software Engineer"
	)
	fmt.Printf("Employed: %t, Salary: $%d, Position: %s\n", isEmployed, salary, position)

	const typedAge int = 38
	const untypedAge = 38 
	fmt.Println(typedAge == untypedAge) // true, same value but different types

	// no enums in Go, but we can use constants
	const (
		Jan int = iota + 1 // iota starts at 0, so Jan is 1
		Feb // 2
		Mar // 3
		Apr // 4
	)
}
