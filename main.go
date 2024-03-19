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

	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)
	word, err := dictionary.Search(baseWord)
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println(word)
	}

}