package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("Welcome to Structure")

	p1 := Person{
		Name: "Omshree",
		Age:  23,
	}
	fmt.Printf("Structure data is %+v", p1)

}
