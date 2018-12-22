package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	iceCream  []string
}

func main() {
	p1 := person{
		firstName: "Jeff",
		lastName:  "Bridges",
		iceCream:  []string{"vanilla", "strawberry"},
	}
	p2 := person{
		firstName: "Dick",
		lastName:  "Hed",
		iceCream:  []string{"chocolate", "coffee", "neopolitan", "rum n raison"},
	}
	fmt.Println(p1.firstName, p1.lastName)
	for i, v := range p1.iceCream {
		fmt.Println(i, v)
	}
	fmt.Println(p2.firstName, p2.lastName)
	for i, v := range p2.iceCream {
		fmt.Println(i, v)
	}
	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
