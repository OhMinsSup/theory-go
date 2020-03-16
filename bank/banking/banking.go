package banking

import (
	"errors"
	"fmt"
)

// Account Struct
type account struct {
	owner   string
	balance int
}

// Deposit x amount on your account
func (a *account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account
func (a account) Balance() int {
	return a.balance
}

// Withdraw x amount on your account
func (a *account) Withdraw(amount int) error {
	if a.balance < amount {
		return errors.New("Can't withraw you are porr")
	}
	a.balance -= amount
	return nil
}

// ChangeOwner of the account
func (a *account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner of the account
func (a account) Owner() string {
	return a.owner
}

func (a account) String() string {
	return fmt.Sprint(a.Owner(), "'s account. \nHas: ", a.Balance())
}

// NewAccount creates Account
func NewAccount(owner string) *account {
	a := account{owner: owner, balance: 0}
	return &a
}
