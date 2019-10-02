package main

import "fmt"

func main() {
	//long method
	i := 1
	for i <= 10 {
		fmt.Println(i)
		//i = i + 1
		i++
	}

	//short method
	for i := 1; i <= 10; i++ {
		fmt.Printf("Number %d\n", i)
	}

	//fizzbuzz interview question
	for i := 0; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Printf("FizzBuzz number was %d\n", i)
		} else if i%3 == 0 {
			fmt.Printf("Fizz number was %d\n", i)
		} else if i%5 == 0 {
			fmt.Printf("Buzz number was %d\n", i)
		} else {
			fmt.Println(i)
		}

	}
}
