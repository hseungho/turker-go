package main

import "fmt"

func hasRichFriend() bool {
	return true
}

func getFriendsCount() int {
	return 3
}

func main() {
	price := 35000

	if price > 50000 {
		if hasRichFriend() {
			fmt.Println("앗 신발끈이 풀렸네")
		} else {
			fmt.Println("나눠내자")
		}
	} else if price >= 30000 && price <= 50000 {
		if getFriendsCount() > 3 {
			fmt.Println("어이쿠 신발끈이...")
		} else {
			fmt.Println("사람도 얼마 없는데 나눠내자")
		}
	} else {
		fmt.Println("오늘은 내가 쏜다")
	}
}
