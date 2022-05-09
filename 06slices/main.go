package main

import "fmt"

func main() {
	fmt.Println("Welcome to slices")

	var fruitList = []string{"Apple", "Orange", "Peach"}

	fmt.Println(fruitList)

	fruitList = append(fruitList, "Mango")

	fmt.Println(fruitList)

	fruitList = append(fruitList[:2])
	fmt.Println(fruitList)
}
