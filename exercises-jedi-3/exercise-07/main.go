package main

import "fmt"

func main() {
	name := "Fred"
	if name == "James" {
		fmt.Println(name)
	} else if name == "Jimmy" {
		fmt.Println("Not James Bond!")
	} else {
		fmt.Println("Dunno who")
	}
}
