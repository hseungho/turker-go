package main

import "fmt"

// 타입 변환.
// int16 -> int8로 변환 시, int8 범위를 벗어나는 경우
// 3456이 int8로 변환 시, 상위 1바이트가 삭제되기 때문에, -128이 됨
func main() {
	var a int16 = 3456
	var c int8 = int8(a)

	fmt.Println(a)
	fmt.Println(c)
}
