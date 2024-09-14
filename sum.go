package main

import "fmt"

func sum() {
	var (
		rng    int
		result int
	)

	fmt.Println("Enter range")

	fmt.Scan(&rng)

	for i := 0; i <= rng; i++ {
		result += i
	}

	fmt.Printf("Sum of all digits in the given range (%v) = %v", rng, result)
}
