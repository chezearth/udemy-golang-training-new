package main

import "fmt"

func main() {
	x := true
	fmt.Println(x && true)
	fmt.Println(x && false)
	fmt.Println(x || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}
