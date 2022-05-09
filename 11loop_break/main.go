package main

import "fmt"

func main() {
	fmt.Println("Welcome to Loop in go")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thrusday", "Friday", "Saturday"}

	for k := range days {
		fmt.Println(days[k])
	}
	fmt.Println("--------------------------------")
	for _, val := range days {
		fmt.Println(val)
	}
}
