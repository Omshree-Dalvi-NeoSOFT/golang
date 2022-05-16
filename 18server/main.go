package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to creating server")

	PerformGetRequest()
}

func PerformGetRequest() {
	const myURL = "https://www.thunderclient.com/welcome"

	response, err := http.Get(myURL)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status Code :", response.StatusCode)
	fmt.Println("Content Length :", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)

	var responseString strings.Builder

	byteContent, _ := responseString.Write(content)

	fmt.Println("ByteCount is :", byteContent)
	fmt.Println(responseString.String())
	//fmt.Println(string(content))

}
