package accounts

// Account struct (Account로 시작하는 코멘트가 있어야 error가 뜨지 않는다.)
type Account struct {
	owner   string // 소문자로 되어있기 때문에 private struct
	balance int
}

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}
