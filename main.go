package main

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int { //a와 b의 데이터타입이 int이고, 리턴값의 데이터타입이 int이다.
	return a * b
}

// 2개의 리턴이 가능하다!!
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	// var name string = "hayeon"
	name := "hayeon" // 위에거랑 같음 but short. 여기서 처음 설정한 타입이 계속 적용된다.(String으로 지정했다면 끝까지 String)
	name = "Dana"
	fmt.Println(name)

	fmt.Println(multiply(2, 2))

	totalLength, upperName := lenAndUpper("hayeon")
	fmt.Println(totalLength, upperName)

	repeatMe("Hayeon", "Dana", "Harley", "Mike")
}
