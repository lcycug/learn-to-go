package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	lastName  string
	firstName string
	contactInfo
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(firstName string) {
	(*pointerToPerson).firstName = firstName
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Harrison",
		contactInfo: contactInfo{
			zipCode: 123456,
			email:   "test@test.com",
		},
	}
	alex.updateName("Alexander")
	alex.print()
}
