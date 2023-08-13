package main

import "fmt"

func main() {
	str := "Hello\tGo\t\tWorld\n\"Go\"is Awesome!\n"

	fmt.Printf(str)
	fmt.Printf("%s", str)
	fmt.Printf("%q", str)
}
