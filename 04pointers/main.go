package main

import "fmt"

func main() {
	fmt.Println("hello pointers")

	myNumber := 31

	ptr := &myNumber

	fmt.Println(*ptr)

	*ptr = *ptr - 18

	fmt.Println(*ptr)
}
