package main

import "fmt"

func odd() {
	var rng int
	fmt.Println("Enter range")

	fmt.Scan(&rng)

	fmt.Println("Following odd numbers lie in the given range:")

	for i := 0; i < rng; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}
