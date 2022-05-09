package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome"

	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your Name")

	val, _ := reader.ReadString('\n')

	fmt.Println(welcome, val)
}
