package main

import "fmt"

const Pig = 1
const Cow = 2
const Chicken = 3

func PrintAnimal(animal int) {
	if animal == Pig {
		fmt.Println("꿀꿀")
	} else if animal == Cow {
		fmt.Println("음메")
	} else if animal == Chicken {
		fmt.Println("꼬끼오")
	} else {
		fmt.Println("...?")
	}
}

func main() {
	PrintAnimal(Pig)
	PrintAnimal(Cow)
	PrintAnimal(Chicken)
	PrintAnimal(4)
	/*
		꿀꿀
		음메
		꼬끼오
		...?
	*/
}
