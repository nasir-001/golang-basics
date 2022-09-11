package main

import "fmt"

func main() {
	ids := []int{1, 2, 4, 5, 6, 7, 8, 9, 0}

	for i, x := range ids {
		fmt.Printf("index %d: number %d\n", i, x)
	}

	sum := 0

	for _, x := range ids {
		sum += x
	}

	fmt.Println(sum)

	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range emails {
		fmt.Println("Name: " + k)
	}
}
