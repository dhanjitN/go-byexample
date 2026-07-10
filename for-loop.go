package main

import "fmt"

func main() {
	i := 1
	fmt.Println("loop with just the condition")
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	fmt.Println("-------------------------------------")
	fmt.Println("Classic for loop with iterator , condition, after")
	for j := 1; j < 3; j++ {
		fmt.Println(j)
	}

	fmt.Println("-------------------------------------")
	fmt.Println("Loop without any condition , inside break")
	end := 10
	for {
		fmt.Println(end)
		end--
		if end == 0 {
			break
		}
	}

	fmt.Println("-------------------------------------")
	fmt.Println("Using Range")
	for n := range 6 {
		// 6 is not included i.e < 6
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
