package main

import "fmt"

var g int = 10

func main() {
	var m int = 20

	{
		var s int = 50
		fmt.Println(m, s, g)
	}

	//m = s + 20 // Error 발생 <- 변수 s를 찾지 못함.
}
