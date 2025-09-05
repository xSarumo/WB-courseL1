package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

func (targ *Point) Distance(other *Point) float64 {
	distX := other.x - targ.x
	distY := other.y - targ.y
	return math.Hypot(distX, distY)
}

func main() {
	pointA := NewPoint(4, 5)
	pointB := NewPoint(3, 2)

	fmt.Println(pointA.Distance(pointB))
}
