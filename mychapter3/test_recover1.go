package main

import (
	"fmt"
)

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Runtime error caught: %v", r)
		}
	}()

	b := 2

	c := 10/b	// Runtime error caught: runtime error: integer divide by zero%

	fmt.Println("c =", c)
}
