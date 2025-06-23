package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	person := Person{Name: "John", Age: 30}

	employee := struct {
		name string
		id   int
	}{
		name: "Alice",
		id:   123,
	}

	type Address struct {
		Street string
		City   string
	}

	type Contact struct {
		Name    string
		Address Address
		Phone   string
	}

	contact := Contact {
		Name: "Mark",
		Address: Address {
			Street: "123 Main St",
			City: "Anytown",
		},
	}

 	fmt.Println("This is our contact", contact);
	fmt.Println("This is our employee", employee);

	fmt.Println("Name before", person.Name);
	person.modifyPersonName("Ben")	
	fmt.Println("Name after", person.Name);

	x := 20
    ptr := &x
	fmt.Printf("Value of x: %d and address of x %p\n", x, ptr);
	*ptr = 30
	fmt.Printf("Value of new x: %d and address of x %p\n", x, ptr);
}

// Name is only modified in the scope of the function
// func modifyPersonName(p *Person) {
// 	p.Name = "Ben"
// 	fmt.Println("Inside Name after", p.Name);
// }

func (p *Person) modifyPersonName(name string) {
	p.Name = name
	fmt.Println("Inside method name", p.Name);
}


