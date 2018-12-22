package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	iceCream  []string
}

func main() {
	person1 := person{
		firstName: "Jeff",
		lastName:  "Bridges",
		iceCream:  []string{"vanilla", "strawberry"},
	}
	person2 := person{
		firstName: "Dick",
		lastName:  "Hed",
		iceCream:  []string{"chocolate", "coffee", "neopolitan", "rum n raison"},
	}
	fmt.Println(person1.firstName, person1.lastName)
	for i, v := range person1.iceCream {
		fmt.Println(i, v)
	}
	fmt.Println(person2.firstName, person2.lastName)
	for i, v := range person2.iceCream {
		fmt.Println(i, v)
	}
}
