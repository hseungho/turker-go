package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	for {
		time.Sleep(time.Second)
		fmt.Println(i)
		i++
	}
	/*
		1
		2
		3
		4
		5
		...
	*/
}
