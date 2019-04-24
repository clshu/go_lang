package main

import "fmt"

// Even Or Odd

func even_or_odd() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, value := range numbers {
		if value%2 == 0 {
			fmt.Printf("%d is an even number\n", value)
		} else {
			fmt.Printf("%d is a odd number\n", value)
		}
	}
}
