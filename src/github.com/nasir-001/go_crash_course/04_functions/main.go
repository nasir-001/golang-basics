package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1, num2 int) int {
	return num1 + num1
}

func main() {
	fmt.Println(getSum(2, 6))
}
