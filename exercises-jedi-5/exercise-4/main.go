package main

import "fmt"

func main() {
	x := struct {
		address  map[string]string
		products []string
	}{
		address:  map[string]string{"number": "19", "street": "Sans Souci", "suburb": "Newlands"},
		products: []string{"beer", "rugby", "cricket"},
	}
	fmt.Println(x)
	fmt.Println(x.address["street"])
	fmt.Println(x.products[0])
}
