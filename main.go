package main

import "fmt"

func canIDrink(age int) bool {
	// switch문
	// switch {
	// case age < 18:
	// 	return false
	// case age == 18:
	// 	return true
	// }
	// return false

	// 변수를 생성하는 switch문
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

func main() {
	fmt.Println(canIDrink(16))
}
