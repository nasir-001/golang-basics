package main

import "fmt"

func main() {
	// var name string = "Nasir"
	name, email := "Nasir", "nasirlawal001@gmail.com"
	var age int32 = 32
	var isCool = true
	isCool = false
	size := 3.12

	fmt.Println(name, age, isCool, size, email)
	fmt.Printf("%T\n", size)
}
