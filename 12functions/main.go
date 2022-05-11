package main

import "fmt"

func main() {
	fmt.Println("Welcome to Functions")

	result := proAdder(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	fmt.Println("Addition is ", result)
}

func proAdder(values ...int) int {
	total := 0

	for _, val := range values {
		total = total + val
	}
	return total
}
