package main

import "fmt"

const PI = 3.14
const FloatPI float64 = 3.14

func main() {
	var a int = PI * 100
	//var b int = FloatPI * 100 -> 컴파일 에러
	//var c int = int(FloatPI) * 100 -> 컴파일 에러

	fmt.Println(a)
	/*
		314
	*/
}
