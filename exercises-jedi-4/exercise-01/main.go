package main

import (
	"fmt"
)

func main() {

	x := [5]int{23, 17, 29, 31, 7}
	fmt.Println(x)
	for i, v := range x {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", x)

}
