package main

import "fmt"

func greeting(name string) string {
	return "hello" + name
}

func getSum(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(greeting("Toby"))
	fmt.Println(getSum(25, 26))
}
