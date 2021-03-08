package main

import "fmt"

func superAdd(numbers ...int) int {
	// for문 기본형
	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println(numbers[i])
	// }
	// return 1

	// range는 index를 사용함
	// for index, number := range numbers {
	// 	fmt.Println(index, number)
	// }
	// return 1

	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)
}
