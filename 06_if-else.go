package main

import "fmt"

func main() {
	// Odd even check
	if 8%2 == 0 {
		fmt.Println("8 is even ")
	} else {
		fmt.Println("8 is odd ")
	}

	// Logical operations
	if 8%2 == 0 && 4%2 == 0 {
		fmt.Println("Both 8 and 2 are even ")
	}
	if 7%2 == 0 || 8%2 == 0 {
		fmt.Println("Either 7 or 8 is even ")
	}

	// can start with a statments , every conditional has access to this statement
	if num := 200; num < 0 {
		fmt.Println(num, "is negative")
	} else if num > 10 && num < 100 {
		fmt.Println(num, "has 2 digits ")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits ")
	}

	// fmt.Println(num, "is accessible here") // not accessible

	// Go has no ternay operator

}
