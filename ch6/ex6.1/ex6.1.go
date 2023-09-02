package main

import "fmt"

func main() {
	var x int32 = 7
	var y int32 = 3

	var s float32 = 3.14
	var t float32 = 5

	fmt.Printf("x + y = %d\n", x+y)
	fmt.Printf("x - y = %d\n", x-y)
	fmt.Printf("x * y = %d\n", x*y)
	fmt.Printf("x / y = %d\n", x/y)
	fmt.Printf("x %% y = %d\n", x%y)

	fmt.Printf("s * t = %g\n", s*t)
	fmt.Printf("s / t = %g\n", s/t)
}
