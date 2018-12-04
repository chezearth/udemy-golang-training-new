package main

import "fmt"

func main() {
	name := "James"
	if name == "James" {
		fmt.Println("Bond")
	}
	if name == "Roger" {
		fmt.Println("Moore")
	}
	if name == "Sean" || name == "James" {
		fmt.Println("Connery as Bond")
	}
}
