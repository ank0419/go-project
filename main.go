package main

import (
	"fmt"
	"log"

	"github.com/ahk0419/go-project/accounts"
)

func main() {
	account := accounts.NewAccount("ahk0419")
	account.Deposit(10)

	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(account.Balance(), account.Owner())

	dictionary := dic

}