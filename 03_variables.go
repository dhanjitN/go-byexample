package main

import "fmt"

func main() {
	var name string = "Paul"
	fmt.Println(name)

	var a = "Joe" // automatic type infertion
	a = "Willam"
	fmt.Println(a)

	var i, j int = 1, 2 // multiple variable declaration
	fmt.Println(i, j)

	var e int
	fmt.Println("e", e) // zero valued

	var isGood bool
	fmt.Println("isGood", isGood) // false

	age := 18
	fmt.Println("age", age) // shorthand for variable initialistaion and declaration
	// this syntax is available only inside functions !
}
