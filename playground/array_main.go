package main

import "fmt"

func main() {
	names := []string{"hayeon", "harley", "dana"} // 길이에 제한이 없는 Array
	fmt.Println(names)
	names = append(names, "mike") // index를 지정해서 추가할 수는 없고, 반드시 append를 사용한다.
	fmt.Println(names)
}
