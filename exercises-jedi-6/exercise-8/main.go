package main

import "fmt"

func main() {

	y := foo()
	fmt.Printf("%T\n", y)
	fmt.Println(y())

}

func g() string {
	x := "Hello this is g()"
	fmt.Printf("%T\n", x)
	fmt.Printf("%s\n", x)
	return x
}

func foo() func() string {
	return g
}
