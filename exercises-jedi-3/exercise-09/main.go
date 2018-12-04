package main

import "fmt"

func main() {
	var favSport string = "running"
	switch favSport {
	case "golf":
		fmt.Printf("Favourite Sport is %v.\n", favSport)
	case "running":
		fmt.Printf("Makes me tired: %v!\n", favSport)
	case "rugby":
		fmt.Printf("Rugger! (%v)\n", favSport)
	}
}
