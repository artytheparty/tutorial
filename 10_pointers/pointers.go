package main

import "fmt"

func main() {
	a := 5
	//assigning b to pointer of a
	//will actually print the memory address
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", b)

	//use * to read values from address
	fmt.Println(*b)
	fmt.Println(*&a)

	//change val with pointer
	*b = 10
	fmt.Println(a) //eaqual to 10 ?
}
