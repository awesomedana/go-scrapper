package main

import (
	"fmt"

	"github.com/awesomedana/go-scrapper/accounts"
)

func main() {
	account := accounts.NewAccount("dana")
	fmt.Println(account)
}
