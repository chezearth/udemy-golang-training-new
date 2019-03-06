package main

import "fmt"

func main() {
	defer foo()
	bar()
}

func foo() {
	fmt.Println(42)
}

func bar() {
	fmt.Println("Hello world")
}
