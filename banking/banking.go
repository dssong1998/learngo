package banking

import (
	"errors"
	"fmt"
)

// Account struct
type Account  struct {
	owner string
	balance int
}
// error message
var errNoMoney = errors.New("Not enough balance in your account")

// NewAccount create account
func NewAccount(owner string) *Account{
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit on account
func (a *Account) Deposit(amount int){
	a.balance += amount
}

// Balance show
func (a Account) Balance() int {
	return a.balance
}
// Owner show
func (a Account) Owner() string {
	return a.owner
}

// Withdraw from account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -=amount
	return nil
} 

// ChangeOwner of account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
} 

func (a Account) String () string {
	return fmt.Sprint(a.Owner(),"'s account\nHas: $", a.Balance())
}