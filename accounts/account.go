package accounts

import (
	"errors"
	"fmt"
)

// Account is for user (private fields)
type Account struct {
	owner   string
	balance int
}

// errFoo is convention
var (
	errNoMoney = errors.New("Can't withdraw: you don't have that money")
)

// NewAccount is factory to make Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	// 새로 생성한 account를 value로 copy 시키고 싶지 않기 때문에 &로 전달
	return &account
}

/*
Deposit amount money on your account
receiver의 conv는 struct의 앞글자 소문자
`Pointer receiver`: go 에서는 value로 전달하기 때문에 copy가 일어난다. 이를 방지하기 위해서 `*receiver`를 사용한다.
*/
func (a *Account) Deposit(ammount int) {
	defer a.Balance()
	a.balance += ammount
}

// Balance shows balance of a account
func (a *Account) Balance() int {
	// receiver 복사해도 상관없다.
	fmt.Printf("%s has $%d\n", a.owner, a.balance)
	return a.balance
}

// Withdraw amount from a account
func (a Account) Withdraw(amount int) error {
	defer a.Balance()
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil // error의 nil type
}
