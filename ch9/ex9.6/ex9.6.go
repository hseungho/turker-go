package main

import "fmt"

func getMyAge() (int, bool) {
	return 33, true
}

func main() {
	if age, ok := getMyAge(); ok && age < 20 {
		fmt.Println("You are young", age)
	} else if ok && age < 30 {
		fmt.Println("Nice age", age)
	} else if ok {
		fmt.Println("You are beautiful", age)
	} else {
		fmt.Println("Error")
	}

	//fmt.Println("Your age is", age) -> 컴파일 에러 - age 변수는 이미 소멸되었음

	/*
		You are beautiful 33
	*/
}
