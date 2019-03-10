package main

import "fmt"

func main() {

	x := func(y string) string {
		// return 10
		return "Hello " + y
	}

	fmt.Println(x("Charles"))
	fmt.Printf("%T\n", x)

}
