package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner   string
	balance int
}

//누구나 나의 Balance를 0 마음대로 조정하지 못하도록 대문자를 소문자로 지정후 퍼블릭을 만든다
//실제 메모리 account를 return
// NewAccount create Account
func NewAccount(owner string) *Account{
	account := Account{owner: owner, balance: 0}
	return &account
}

//balance ++ 증가 /-- 감소
//Deposit x amount onyour account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errors.New("No, Money")
	}
	a.balance -= amount
	return nil
}

//Balance of account
func (a Account) Balance() int {
	return a.balance
}

//ChangeOwner of the Account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account. \nHas: ", a.Balance())
}