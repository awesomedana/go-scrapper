package main

import "fmt"

func main() {
	a := 2
	b := &a
	fmt.Println(a, b)
	fmt.Println(&a, b) // 메모리 주소를 출력
	a = 5
	fmt.Println(*b) // 메모리 주소에 담긴 값을 출력
	*b = 20
	fmt.Println(a, b) // *b는 a를 볼 수 있는 Pointer이기 때문에 a의 값을 변경할 수 있다!
}
