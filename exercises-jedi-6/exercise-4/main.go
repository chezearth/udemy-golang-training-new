package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   33,
	}
	p1.speak()
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, "and my age is", p.age)
}
