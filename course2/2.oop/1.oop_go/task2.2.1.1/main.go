package main

import (
	"fmt"
)

type Accounter interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	Balance() float64
}

type CurrentAccount struct {
	balance float64
}

type SavingsAccount struct {
	balance float64
}

func (c *CurrentAccount) Deposit(amount float64) error {
	if amount < 0 {
		return fmt.Errorf("can't deposit negative amount")
	}
	c.balance += amount
	return nil
}

func (c *SavingsAccount) Deposit(amount float64) error {
	if amount < 0 {
		return fmt.Errorf("can't deposit negative amount")
	}
	c.balance += amount
	return nil
}

func (c *CurrentAccount) Withdraw(amount float64) error {
	if amount < 0 {
		return fmt.Errorf("can't withdraw negative amount")
	}
	if c.balance < amount {
		return fmt.Errorf("not enough balance to withdraw. Balance: %.2f, amount: %.2f", c.balance, amount)
	}
	c.balance -= amount
	return nil
}

func (c *SavingsAccount) Withdraw(amount float64) error {
	if amount < 0 {
		return fmt.Errorf("can't withdraw negative amount")
	}
	if c.balance < amount {
		return fmt.Errorf("not enough balance to withdraw. Balance: %.2f, amount: %.2f", c.balance, amount)
	}
	if c.balance < 500 {
		return fmt.Errorf("can't withdraw from savings account with balance less than 500. Current balance: %.2f", c.balance)
	}
	c.balance -= amount
	return nil
}

func (c *CurrentAccount) Balance() float64 {
	return c.balance
}

func (c *SavingsAccount) Balance() float64 {
	return c.balance
}

func ProcessAccount(a Accounter) {
	a.Deposit(500)
	a.Withdraw(200)
	fmt.Printf("Balance: %.2f\n", a.Balance())
}

func main() {
	c := &CurrentAccount{}
	s := &SavingsAccount{}
	ProcessAccount(c)
	ProcessAccount(s)
}
