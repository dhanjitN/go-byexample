package main

import (
	"fmt"
	"math"
)

const owner_name string = "Dhanjit"

func main() {
	fmt.Println("Owner :", owner_name)

	const owner_age int = 20
	fmt.Println("Owner's age:", owner_age)

	const n = 50000
	fmt.Println("n/10", n/10)
	fmt.Println("n/9.9", n/9.9)
	fmt.Println("sin(n)", math.Sin(n))
	// if a integer constant is not given a type , it get's converted to the type that is required

}
