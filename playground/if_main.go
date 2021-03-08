package main

import "fmt"

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 { // if문의 조건문 안에서 변수를 만들 수 있다! (단, if문 안에서만 사용한다)
		return false
	}
	return true
}

func main() {
	fmt.Println(canIDrink(16))
}
