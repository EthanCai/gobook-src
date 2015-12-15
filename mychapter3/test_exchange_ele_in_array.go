package main

import (
	"fmt"
)

func main() {

	a1 := []int { 1, 2, 3 }

	a1[0], a1[1] = a1[1], a1[0]
	a1[0], a1[0] = a1[0], a1[0]

	fmt.Println(a1)
}
