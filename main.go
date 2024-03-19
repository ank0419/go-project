package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/ahk0419/go-project/accounts"
	"github.com/ahk0419/go-project/mydict"
)

var errRequestFailed = errors.New("Request Failed")

type resultRes struct {
	url string
	status string
}

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

	results := make(map[string]string)
	c := make(chan resultRes)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}


	for _, url := range urls {
		go hitURL(url, c)
	}

	for i :=0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}
	for url, status := range results {
		fmt.Println(url, status)
	}
}

func hitURL(url string, c chan resultRes) {
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	} 
	c <- resultRes{url: url, status: status}
}