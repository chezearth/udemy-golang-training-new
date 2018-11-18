package main

import "fmt"

// declare type and value
var x = 42
var y = "James Bond"
var z = true

func main() {
	fmt.Printf("%v\n", x)
	fmt.Printf("%v\n", y)
	fmt.Printf("%v\n", z)
	// ZERO values if only type is used in package scope

	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)

}
