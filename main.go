package main

import (
	"fmt"

	"github.com/ahk0419/go-project/accounts"
)

func main() {
	account := accounts.NewAccount("ahk0419")
	account.Deposit(10)

	fmt.Println(account.Balance())
	account.Withdraw(20)
	fmt.Println(account.Balance())

}