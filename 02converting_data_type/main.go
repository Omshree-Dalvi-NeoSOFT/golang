package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Welcome to our application ........")

	fmt.Println("Please provide 1 out of 5 ratings....")

	reader := bufio.NewReader(os.Stdin)

	value, _ := reader.ReadString('\n')

	fmt.Println("Thank-you, you have rated ", value)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(value), 64)

	if err != nil {
		fmt.Println("You got an error !!!!...", err)
	} else {
		fmt.Println("We have added + 1 to your ratings,", numRating+1)
	}
}
