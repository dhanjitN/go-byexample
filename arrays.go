package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("arr: ", a)
	// by default values are 0 and garbagge values

	a[4] = 100
	fmt.Println("set: ", a)
	fmt.Println("get: ", a[4])

	fmt.Println("len: ", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("arr b :", b)

	// ... means fiding the finding the number of elements in the array if you don't know it already
	b = [...]int{2, 4, 5, 6, 7}
	fmt.Println("arr b:", b)

}
