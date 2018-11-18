package main

import "fmt"

// declare type only
var x int
var y string
var z bool

func main() {
	fmt.Printf("%v\n", x)
	fmt.Printf("%v\n", y)
	fmt.Printf("%v\n", z)
	// ZERO values if only type is used in package scope

}
