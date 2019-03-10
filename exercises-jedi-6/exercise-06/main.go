package main

import "fmt"

func main() {
	x := 21
	fmt.Println("x =", x)
	func(x int) {
		x++
		fmt.Println("add one to x:", x)
	}(x)
	fmt.Println("x (original scope):", x)
}
