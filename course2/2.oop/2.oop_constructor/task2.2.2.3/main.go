package main

import (
	"fmt"
	"sync"
)

type Account interface {
	Deposit(amount float64)
	Withdraw(amount float64) error
	Balance() float64
}

type Customer struct {
	ID   int
	Name string
	Account
}

type SavingsAccount struct {
	balance float64
	mutex   sync.Mutex
}

type CustomerOption func(*Customer)

func NewCustomer(ID int, cops ...CustomerOption) *Customer {
	customer := &Customer{ID: 1}
	for _, cop := range cops {
		cop(customer)
	}
	return customer
}

func WithName(name string) CustomerOption {
	return func(c *Customer) {
		c.Name = name
	}
}

func WithAccount(account Account) CustomerOption {
	return func(c *Customer) {
		c.Account = account
	}
}

func (c *SavingsAccount) Deposit(amount float64) {
	c.mutex.Lock()
	c.balance += amount
	c.mutex.Unlock()
}

func (c *SavingsAccount) Withdraw(amount float64) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if amount < 0 {
		return fmt.Errorf("can't withdraw negative amount")
	}
	if c.balance < amount {
		return fmt.Errorf("not enough balance to withdraw. Balance: %.2f, amount: %.2f", c.balance, amount)
	}
	if c.balance < 1000 {
		return fmt.Errorf("can't withdraw from savings account with balance less than 1000. Current balance: %.2f", c.balance)
	}
	c.balance -= amount
	return nil
}

func (c *SavingsAccount) Balance() float64 {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.balance
}

func main() {
	savings := &SavingsAccount{}
	savings.Deposit(1000)

	customer := NewCustomer(1, WithName("John Doe"), WithAccount(savings))

	err := customer.Account.Withdraw(100)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Customer: %v, Balance: %v\n", customer.Name, customer.Account.Balance())
}
