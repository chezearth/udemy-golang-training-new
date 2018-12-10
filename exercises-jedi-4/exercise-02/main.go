package main

import "fmt"

func main() {

	x := []int{1, 2, 3, 5, 7, 11, 13, 17, 19, 23}
	fmt.Printf("%T\t%v\n", x, x)

	for i, v := range x {
		fmt.Printf("%v\t%T\t%v\n", i, v, v)
	}

}
