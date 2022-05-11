package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Status bool
	Email  string
}

func main() {
	fmt.Println("Welcome to Structure")

	p1 := Person{
		Name:   "Omshree",
		Age:    23,
		Status: true,
		Email:  "Omshree.dalvi@neosoftmail.com",
	}
	fmt.Printf("Structure data is %+v \n", p1)
	p1.getStatus()
	p1.newEmail()

}

func (p Person) getStatus() {
	fmt.Println("The status of user is", p.Status)
}

func (p Person) newEmail() {
	p.Email = "newEmail@neo.co.in"
	fmt.Println("updated Email of user is ", p.Email)
}
