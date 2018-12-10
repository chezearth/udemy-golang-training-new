package main

import "fmt"

func main() {

	bondJames := []string{`Shaken, not stirred`, `Martinis`, `Women`}
	moneypennyMiss := []string{`James Bond`, `Literature`, `Computer Science`}
	noDr := []string{`Being evil`, `Ice cream`, `Sunsets`}

	m := map[string][]string{
		"bond_james":      bondJames,
		"moneypenny_miss": moneypennyMiss,
		"no_dr":           noDr,
	}

	for k, v := range m {
		for i, val := range v {
			fmt.Printf("key: %v\t index: %v\t value: %v\n", k, i, val)
		}
	}

	fmt.Println("Add in a record")
	m["q"] = []string{"gadgets", "inventing", "technology"}

	for k, v := range m {
		for i, val := range v {
			fmt.Printf("key: %v\t index: %v\t value: %v\n", k, i, val)
		}
	}

}
