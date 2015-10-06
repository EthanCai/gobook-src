package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x float64
	y float64
	r float64
}

func circleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}

func doubleR(c Circle) {
	c.r = 2 * c.r
}

func circleArea2(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

func doubleR2(c *Circle) {
	c.r = 2 * c.r
}

func main() {
	circle1 := Circle{0, 0, 5}
	fmt.Println(circleArea(circle1))

	doubleR(circle1)
	fmt.Println(circleArea(circle1))

	circle2 := Circle{0, 0, 5}
	fmt.Println(circleArea(circle2))

	doubleR2(&circle2)
	fmt.Println(circleArea(circle2))
}
