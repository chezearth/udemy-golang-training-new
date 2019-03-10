package main

import "fmt"

func main() {
	x := 40
	y := 2
	fmt.Printf("%T\n", x+y)

	foo(func() int {
		return 42
	})

}

func foo(bar func() int) int {
	fmt.Println(bar())
	return bar()
}
