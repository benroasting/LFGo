package main

import (
	"fmt"
)

func main() {
	age := 30
	if age >= 18 {
		fmt.Println("You are an adult")
	} else if age >= 13 {
		fmt.Println("You are a teenager")
	} else {
		fmt.Println("You are a child")
	}

	day := "Tuesday"
	switch day {
	case "Monday":
		fmt.Println("Start of the week")
	case "Tuesday", "Wednesday", "Thursday":
		fmt.Println("Midweek")
	case "Friday":
		fmt.Println("TGIF")
	default: 
		fmt.Println("Weekend") 
	}

	for i := 0; i < 5; i++ {
		fmt.Println("Iteration:", i)
	}

	// while loop in Go is just a for loop with a condition
	counter := 0
	for counter < 5 {
		fmt.Println("Counter:", counter)
		counter++
	}

	// infinite loop if needed
	
	iterations := 0
	for  {
		if iterations >3 {
			break
		}
		iterations++
	}

	// Arrays and slices
	// Arrays are fixed size, slices are dynamic
	// [5]int is an array of 5 integers
	number := [5]int{10,20,30,40,50}
	// arrays cannot be changed after declaration
	fmt.Printf("this is our array: %v\n", number)
	fmt.Println("this is our last value", number[len(number)-2])

	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Printf("this is our matrix: %v\n", matrix)

	// allNumbers := number[:];
	// firstThree := number[0:3]

	fruits := []string{"apple", "banana", "cherry"}
	fmt.Printf("this is our slice of fruits: %v\n", fruits)

	fruits = append(fruits, "strawberry")
	fmt.Printf("after appending, our slice of fruits is: %v\n", fruits)

	for index, value := range fruits {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}
}


