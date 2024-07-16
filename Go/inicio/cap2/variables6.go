package cap2

import "fmt"

var a int
var b float32
var c string

func Example6() {
	x = 10
	fmt.Println(x)

	x = 10
	fmt.Println(x)

	// Valor Zero:
	// - int: 0
	// - floats: 0.0
	// - booleans: false
	// - string: ""
	// - pointers, functions, interfaces, slices, channels, maps: nil

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
