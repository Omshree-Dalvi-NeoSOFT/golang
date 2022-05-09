package main

import "fmt"

func main() {
	fmt.Println("Welcome to Array")

	var num [5]int

	num[0] = 1
	num[1] = 11
	num[2] = 111
	num[3] = 1111
	num[4] = 11111

	fmt.Println("Number List is ", num)
	fmt.Println("Number List is ", len(num))

	var alpha = [4]string{"a", "b", "c", "d"}

	fmt.Println("Alphabt List is ", alpha)
	fmt.Println("Alphabt List is ", len(alpha))
}
