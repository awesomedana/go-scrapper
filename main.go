package main

import (
	"fmt"

	"github.com/awesomedana/go-scrapper/accounts"
)

func main() {
	account := accounts.NewAccount("dana")
	account.Deposit(10)
	fmt.Println(account.Balance())
	// err := account.Withdraw(2)
	// if err != nil {
	// 	log.Fatalln(err) // println and kill the program
	// }
	fmt.Println(account.Balance(), account.Owner())
	fmt.Println(account)
}
