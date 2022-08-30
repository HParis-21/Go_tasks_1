/*
Разработать программу нахождения расстояния между двумя точками, которые
представлены в виде структуры Point с инкапсулированными параметрами x,y и
конструктором.
*/
package main

import "math"

type Point struct {
	x float64
	y float64
}

// NewPoint returns new Point initialised with x and y
func NewPoint(x, y float64) Point {
	return Point{x, y}
}

// Distance calculates and returns distance to p2
func (p Point) Distance(p2 Point) float64 {
	return math.Sqrt(math.Pow(p2.x-p.x, 2) + math.Pow(p2.y-p.y, 2))
}

func main() {
	p1 := NewPoint(0, 0)
	p2 := NewPoint(10, 20)
	println(p1.Distance(p2))
}
