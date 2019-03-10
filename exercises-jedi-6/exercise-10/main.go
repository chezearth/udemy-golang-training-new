package main

import "fmt"

var x int

func main() {
	x = 10

	fmt.Println(closeOver())
	fmt.Println(x)

	f := foo()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

}

func closeOver() string {
	fmt.Println(x)
	x := "Fred"
	fmt.Println(x)
	return x
}

func foo() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
