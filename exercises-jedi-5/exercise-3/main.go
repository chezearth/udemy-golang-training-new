package main

import "fmt"

type vehicle struct {
	doors  int
	colour string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t := truck{
		vehicle: vehicle{
			doors:  2,
			colour: "blue",
		},
		fourWheel: true,
	}
	s := sedan{
		vehicle: vehicle{
			doors:  4,
			colour: "red",
		},
		luxury: true,
	}
	fmt.Println(t, s)
	fmt.Println("truck colour: ", t.colour)
	fmt.Println("sedan luxury: ", s.luxury)
}
