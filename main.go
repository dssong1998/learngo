package main

import (
	"fmt"

	"github.com/dssong1998/learngo/banking"
)

func main() {
	account:=banking.NewAccount("DS")
	account.Deposit(10000)
	fmt.Println(account)
	account.Withdraw(3000)
	fmt.Println(account)
}