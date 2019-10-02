package main

import "fmt"

func main() {
	//define maps
	//emails := make(map[string]string)
	//assign key values
	//	emails["bob"] = "bob@gmail.com"
	//	emails["nick"] = "nick@gmail.com"
	//	emails["lucy"] = "blucy@gmail.com"

	emails := map[string]string{"bob": "bob@hotmail.com", "mike": "mike@yahoo.com", "sharon": "sharon@revatourism.org"}

	//prints out datatype and all value/key pairing
	fmt.Println(emails)
	//length of map
	fmt.Println(len(emails))
	//get specific data from key
	fmt.Println(emails["bob"])

	//delete from map
	delete(emails, "bob")
	fmt.Println(emails)

}
