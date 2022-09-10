package main

import "fmt"

func main() {
	// arrays
	// var fruitArr [2]string

	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	fruitArr := [2]string{"Apple", "Orange"}

	fruitSlices := []string{"Apple", "Orange", "Mango"}

	fmt.Println(fruitArr)
	fmt.Println(fruitSlices, len(fruitSlices))
}
