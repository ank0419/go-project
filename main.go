package main

import (
	"fmt"
	"log"

	"github.com/ahk0419/go-project/accounts"
	"github.com/ahk0419/go-project/mydict"
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

	dictionary := mydict.Dictionary{"first": "First word"}
	dictionary.Add()
	definition, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
}