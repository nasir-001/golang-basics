package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// Greeting method (value receiver)
func (self Person) greet() string {
	return "Hello, my name is " + self.firstName + " " + self.lastName + " and I'm " + strconv.Itoa(self.age)
}

// has birthday method (pointer receive)
func (self *Person) hasBirthday() {
	self.age++
}

// getMarried (pointer receiver)
func (self *Person) getMarried(spouseLastName string) {
	if self.gender == "m" {
		return
	} else {
		self.lastName = spouseLastName
	}
}

func main() {
	person1 := Person{firstName: "Aisha", lastName: "Yusuf", city: "Kaduna", age: 20, gender: "f"}

	// fmt.Println(person1)
	// fmt.Println(person1.firstName)
	person1.hasBirthday()
	person1.getMarried("Nasir")
	fmt.Println(person1.greet())

}
