package main

import "fmt"

func main() {
	dan := 2
	b := 1
	for {
		for {
			fmt.Printf("%d * %d = %d\n", dan, b, dan*b)
			b++

			if b == 10 {
				break
			}
		}
		b = 1
		dan++
		if dan == 10 {
			break
		}
	}
	/*
		2 * 1 = 2
		2 * 2 = 4
		2 * 3 = 6
		2 * 4 = 8
		2 * 5 = 10
		2 * 6 = 12
		2 * 7 = 14
		2 * 8 = 16
		2 * 9 = 18
		...중략...
		9 * 1 = 9
		9 * 2 = 18
		9 * 3 = 27
		9 * 4 = 36
		9 * 5 = 45
		9 * 6 = 54
		9 * 7 = 63
		9 * 8 = 72
		9 * 9 = 81
	*/
}
