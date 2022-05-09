package main

import "fmt"

func main() {
	fmt.Println("Welcome to maps")

	names := make(map[string]string)

	names["od"] = "Omshree Dalvi"
	names["ss"] = "Sandeep Sabu"
	names["ns"] = "Neer Shah"
	names["vs"] = "Vivek Sakpal"
	names["rp"] = "Revati Palande"

	fmt.Println(names)
	fmt.Println("abbrivation for od ", names["od"])

	delete(names, "rp")
	fmt.Println(names)

	for key, val := range names {
		fmt.Println("Abbrivation of", key, "is", val)
	}
}
