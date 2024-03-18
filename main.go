package main

import (
	"fmt"

	"github.com/ahk0419/go-project/banking"
)

func main() {
	account := banking.Account{Owner : "ank0419", Balance: 1000}
	fmt.Println(account)
}