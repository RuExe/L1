package main

import (
	"fmt"
	"math"
)

type Point struct {
	x int
	y int
}

func NewPoint(x, y int) Point {
	return Point{
		x: x,
		y: y,
	}
}

func (p *Point) Distance(other Point) float64 {
	x := (other.x - p.x) * (other.x - p.x)
	y := (other.y - p.y) * (other.y - p.y)
	return math.Sqrt(float64(x + y))
}

func main() {
	point := NewPoint(-7, 3)
	point2 := NewPoint(6, 10)
	distance := point.Distance(point2)
	fmt.Println(distance)
}
