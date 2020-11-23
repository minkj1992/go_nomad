package accounts

// Account is for user (private fields)
type Account struct {
	owner   string
	balance int
}

// NewAccount is factory to make Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	// 새로 생성한 account를 value로 copy 시키고 싶지 않기 때문에 &로 전달
	return &account
}
