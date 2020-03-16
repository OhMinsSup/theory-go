package main

import (
	"fmt"
	"log"
	"theory/bank/banking"
)

func main() {
	account := banking.NewAccount("veloss")
	account.Deposit(10)
	err := account.Withdraw(0)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(account)
}
