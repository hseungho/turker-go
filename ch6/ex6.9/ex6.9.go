package main

import (
	"fmt"
	"math/big"
)

func main() {
	a, _ := new(big.Float).SetString("0.1")
	b, _ := new(big.Float).SetString("0.2")
	c, _ := new(big.Float).SetString("0.3")

	d := new(big.Float).Add(a, b)
	fmt.Println(a, b, c, d)
	fmt.Println(c.Cmp(d))
	// x.Cmp(y float64) 함수는 compareTo()와 같이 x가 y보다 작으면 -1, x가 y보다 크면 1, 같으면 0을 반환하는 함수다.
}
