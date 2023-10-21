package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int
	// student   bool

	// Alternative
	firstName, lastName, city, gender string
	age                               int
	student                           bool
}

// Greeting method (value reciever)
func (p Person) greet() string {
	return "Hello, good day. My name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age) + "."
}

// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer reciever)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "M" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// Init person using struct
	// person1 := Person{
	// 	firstName: "Charles",
	// 	lastName:  "Ugbor",
	// 	city:      "Nsk",
	// 	gender:    "M",
	// 	age:       21,
	// 	student:   true,
	// }

	// Alternative
	person1 := Person{
		"Charles",
		"Ugbor",
		"Nsk",
		"M",
		21,
		true,
	}

	person2 := Person{
		"Precious",
		"Chidera",
		"Anambra",
		"F",
		20,
		true,
	}

	// fmt.Println(person1.firstName)
	// person1.age++
	// fmt.Println(person1)

	person1.hasBirthday()

	person1.getMarried("Chidera")
	person2.getMarried("Ugbor")
	fmt.Println(person1.greet())

}
