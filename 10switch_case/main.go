package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to Switch Case")

	rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6) + 1

	//fmt.Println("Dice Number is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("You can open court !!")
		fallthrough
	case 2:
		fmt.Println("you can move 2 spot")
	case 3:
		fmt.Println("you can move 3 spot")
	case 4:
		fmt.Println("you can move 4 spot")
	case 5:
		fmt.Println("you can move 5 spot")
	case 6:
		fmt.Println("you can move 6 spot and roll dice again ")
	default:
		fmt.Println("...")
	}
}
