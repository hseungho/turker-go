package main

import "fmt"

func main() {
	var a = 324.13455
	var c = 3.14

	fmt.Printf("%08.2f\n", a) // 최소 너비 8, 소수점 이하 2자리, 0을 채움 -> 00324.13
	fmt.Printf("%08.2g\n", a) // 최소 너비 8, 소수점 이하 2자리 0을 채움 -> 03.2e+02
	fmt.Printf("%8.5g\n", a)  // 최소 너비 8, 총숫자 5자리 -> 324.13
	fmt.Printf("%f\n", c)     // 소수점 이하 6자리까지 출력 -> 3.140000
}
