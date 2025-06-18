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
	
	result := add(3, 4)
	fmt.Println("This is result", result)

	sum, product := calculateSumAndProduct(10, 10)
	fmt.Printf("Sum: %d, Product: %d\n", sum, product)
}

// if you have a return type, the function must return a value of that type
// if you don't have a return type, the function can return nothing
func add(a int, b int) int {
	return a + b
}

func calculateSumAndProduct(a int, b int) (int, int) {
	return a + b, a * b
}
