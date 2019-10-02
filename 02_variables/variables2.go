package main

import "fmt"

const notCool bool = false

func main() {
	//using var
	var age int = 24
	var isCool bool = true
	//shorthand func declaration
	name := "Khvan"
	//initializing two variables at the same time
	middleName, email := "V", "artkhv15@protonmail.com"

	//print name
	fmt.Println(name+" age:", middleName, age, isCool, notCool)
	fmt.Println(email)
	fmt.Printf("%T\n", isCool)

}
