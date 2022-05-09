package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")

	var fruitList = []string{"Apple", "Orange", "Peach"}

	fmt.Println(fruitList)

	fruitList = append(fruitList, "Mango")

	fmt.Println(fruitList)

	fruitList = append(fruitList[:2])
	fmt.Println(fruitList)

	highScore := make([]int, 4)

	highScore[0] = 89
	highScore[1] = 97
	highScore[2] = 49
	highScore[3] = 79
	//highScore[4] = 69

	fmt.Println(highScore)

	highScore = append(highScore, 69, 22, 74)

	fmt.Println(highScore)

	sort.Ints(highScore)
	fmt.Println(highScore)

	highScore = append(highScore[:3], highScore[4:]...)
	fmt.Println(highScore)

}
