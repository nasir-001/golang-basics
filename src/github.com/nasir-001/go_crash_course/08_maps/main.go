package main

import "fmt"

func main() {
	// emails := make(map[string]string)

	// emails["Bob"] = "bob@gmail.com"
	// emails["Nasir"] = "nasir@gmail.com"
	// emails["mike"] = "mike@gmail.com"

	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}

	fmt.Println(emails)

	delete(emails, "Bob")

	fmt.Println(emails)
}
