package main

import "fmt"

func main() {

	var x = make([][]string, 2)
	x[0] = []string{"James", "Bond", "Shaken, not stirred"}
	x[1] = []string{"Miss", "Moneypenny", "Hellooooo, James"}

	fmt.Println(x)

	for i, vo := range x {
		fmt.Println(i, vo)
		for j, vi := range vo {
			fmt.Println(i, j, vi)
		}
	}

}
