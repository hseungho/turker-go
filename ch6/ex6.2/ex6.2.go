package main

import "fmt"

func main() {
	var x1 int8 = 34   // 0010 0010
	var x2 int16 = 34  // 0000 0000 0010 0010
	var x3 uint8 = 34  // 0010 0010 (부호 없음)
	var x4 uint16 = 34 // 0000 0000 0010 0010 (부호 없음)

	fmt.Printf("^%d = %5d, \t %08b\n", x1, ^x1, uint8(^x1))   // -35, 1101 1101
	fmt.Printf("^%d = %5d, \t %016b\n", x2, ^x2, uint16(^x2)) // -35, 1111 1111 1101 1101
	fmt.Printf("^%d = %5d, \t %08b\n", x3, ^x3, ^x3)          // 221, 1101 1101
	fmt.Printf("^%d = %5d, \t %016b\n", x4, ^x4, ^x4)         // 65501, 1111 1111 1101 1101

	/*
		^34 =   -35, 	 11011101
		^34 =   -35, 	 1111111111011101
		^34 =   221, 	 11011101
		^34 = 65501, 	 1111111111011101
	*/
}
