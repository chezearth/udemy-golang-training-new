package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

type circle struct {
	diameter float64
}

func (sq square) area() float64 {
	return math.Pow(sq.side, 2)
}

func (ci circle) area() float64 {
	return math.Pow(ci.diameter/2, 2) * math.Pi
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
	// switch s.(type) {
	// case square:
	// 	fmt.Println("square:", s.(square).area())
	// case circle:
	// 	fmt.Println("circle:", s.(circle).area())
	// }
}

func main() {
	square1 := square{
		side: 2.2,
	}
	circle1 := circle{
		diameter: 2.2,
	}
	info(square1)
	info(circle1)
}
