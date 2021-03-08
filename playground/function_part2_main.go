package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done.") // defer는 func이 리턴한 후에 실행된다!
	length = len(name)
	uppercase = strings.ToUpper(name)
	return // naked return: 리턴 뒤에 변수를 지정해주지 않아도 된다!
}

func main() {
	totalLength, up := lenAndUpper("hayeon")
	fmt.Println(totalLength, up)
}
