package main

import "fmt"

func getMyAge() int {
	return 22
}

func main() {
	switch age := getMyAge(); true { // true 생략 가능
	case age < 10:
		fmt.Println("Child")
	case age < 20:
		fmt.Println("Teenager")
	case age < 30:
		fmt.Println("20s")
	default:
		fmt.Println("My age is", age)
	}
	/*
		20s
	*/
}
