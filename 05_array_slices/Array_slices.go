package main

import "fmt"

func main() {
	//arrays
	//var fruitArr [2]string
	//assign values
	//fruitArr[0] = "apple"
	//fruitArr[1] = "orange"
	/*
		fruitArr := [2]string{"apple", "orange"}
		//print
		fmt.Println(fruitArr)
		fmt.Println(fruitArr[1])
	*/
	fruitSlice := []string{"apple", "Orange", "grape"}

	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3])
}
