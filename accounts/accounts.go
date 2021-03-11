package accounts

import "errors"

// Account struct (Account로 시작하는 코멘트가 있어야 error가 뜨지 않는다.)
type Account struct {
	owner   string // 소문자로 되어있기 때문에 private struct
	balance int
}

// errNoMoney error
var errNoMoney = errors.New("Can't withdraw")

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on you account(Method)
func (a *Account) Deposit(amount int) { // a Account는 receiver, 항상 struct의 소문자로 시작한다!
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// Withdraw from your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}
