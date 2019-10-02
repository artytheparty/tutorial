package main

import (
	"fmt"
	"strconv"
)

//define person structure
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

//you can do this as well
type Persona struct {
	firstName, lastName, city, gender string
	age                               int
}

//functions for the structure is outside the structure
//Greeting method(value receiver) only took existing values to maybe display or use
func (obj Person) greet() string {
	return "Hello, my name is " + obj.firstName + " " + obj.lastName + " and I am " + strconv.Itoa(obj.age)
}

//hasBirthday method (pointer receive)
func (obj *Person) hasBirthday() {
	obj.age++
}

//getMarried()pointer receiver

func (obj *Person) getMarried(spouseLastName string) {
	if obj.gender == "m" {
		return
	} else {
		obj.lastName = spouseLastName
	}
}

func main() {

	person1 := Person{firstName: "Alan", lastName: "Booker", city: "New Yourk", gender: "m", age: 25}
	person2 := Person{"Alena", "Smith", "Bouston", "f", 30}
	fmt.Println(person1)
	fmt.Println(person2)

	//get a single value from structure
	fmt.Println(person1.firstName)
	person1.age++
	fmt.Println(person1)
	person1.hasBirthday()
	fmt.Println(person1.greet())

	person2.getMarried("Williams")
	fmt.Println(person2)
}
