package main

import "fmt"

func main() {
	ids := []int{33, 75, 54, 21, 1, 43}
	//loop through the ids
	for i, id := range ids {
		fmt.Printf("%d - ID %d\n", i, id)
	}
	//not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	//add all ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("sum", sum)

	//range with map
	emails := map[string]string{"bob": "bob@hotmail.com", "mike": "mike@yahoo.com", "sharon": "sharon@revatourism.org"}
	for k, v := range emails {
		fmt.Printf("%s : %s\n", k, v)
	}
	for k := range emails {
		fmt.Println("name: " + k)
	}
}
