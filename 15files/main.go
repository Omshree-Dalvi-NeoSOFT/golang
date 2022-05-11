package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to Files")

	content := "My name is Omshree Dalvi.."

	file, err := os.Create("./myfile.txt")

	checkNilErr(err)

	length, err := io.WriteString(file, content)

	checkNilErr(err)

	fmt.Println("Length of file is ", length)

	defer file.Close()

	readFile("./myfile.txt")
}

func readFile(fileName string) {
	dataByte, err := ioutil.ReadFile(fileName)
	checkNilErr(err)
	fmt.Println("File Content :- ", string(dataByte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
