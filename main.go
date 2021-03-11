package main

import (
	"fmt"
	"log"

	"github.com/awesomedana/go-scrapper/accounts"
)

func main() {
	account := accounts.NewAccount("dana")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(12)
	if err != nil {
		log.Fatalln(err) // println and kill the program
	}
	fmt.Println(account.Balance())
}
