package main

import (
	"fmt"
	"reflect"
)

// --Output:--
// type: float64
// type: float64
// kind is float64: true
// value: 3.4
func main() {
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))

	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())
}
