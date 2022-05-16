package main

import (
	"fmt"
	"net/url"
)

const myURL string = "https://www.youtube.com:300/watch?v=cl7_ouTMFh0&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=26"

func main() {
	fmt.Println("Welcome to hadling URL")
	//fmt.Println(myURL)

	// Parsing

	result, err := url.Parse(myURL)

	if err != nil {
		panic(err)
	}

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	// fmt.Println(strings.Split(result.RawQuery, "&"))
	fmt.Println(result)

	qparams := result.Query()

	fmt.Println(qparams["index"])

	// construct URL  https://mail.google.com/mail/u/7/#inbox

	partsOfurl := &url.URL{
		Scheme: "https",
		Host:   "mail.google.com",
		Path:   "/mail/u/7/#inbox",
	}

	mainURL := partsOfurl.String()

	fmt.Println(mainURL)
}
