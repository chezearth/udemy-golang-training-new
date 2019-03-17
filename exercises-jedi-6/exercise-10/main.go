package main

import "fmt"

var x int

func main() {

	fmt.Printf("\n%v\n\n\n", "About  C'L'O'S'U'R'E  ...")

	// x = 10
	//
	// fmt.Println("main(): closeOver() method returns =", closeOver())
	// fmt.Println("main(): x =", x)

	a := foo()
	b := foo()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())

}

// func closeOver() string {
// 	fmt.Println("before x initialised locally in closeOver(): x =", x)
// 	x := "Fred"
// 	fmt.Println("after x is initialised and assigned in closeOver(): x =", x)
// 	return x
// }

func foo() func() int {
	y := 0
	return func() int {
		y++
		return y
	}
}
