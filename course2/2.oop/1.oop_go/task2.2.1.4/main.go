package main

import (
	"fmt"
)

type PaymentMethod interface {
	Pay(amount float64) error
}

type CreditCard struct {
	balance float64
}

type Bitcoin struct {
	balance float64
}

func (c *CreditCard) Pay(amount float64) error {
	if amount <= 0.00 {
		return fmt.Errorf("недопустимая сумма платежа")
	}
	if amount > c.balance {
		return fmt.Errorf("недостаточный баланс")
	}
	c.balance -= amount
	fmt.Println("Оплачено", amount, "с помощью кредитной карты")
	return nil
}

func (c *Bitcoin) Pay(amount float64) error {
	if amount <= 0.00 {
		return fmt.Errorf("недопустимая сумма платежа")
	}
	if amount > c.balance {
		return fmt.Errorf("недостаточный баланс")
	}
	c.balance -= amount
	fmt.Println("Оплачено", amount, "с помощью биткоина")
	return nil
}

func ProccessPayment(p PaymentMethod, amount float64) {
	err := p.Pay(amount)
	if err != nil {
		fmt.Println("Не удалось обработать платеж:", err)
	}
}

func main() {
	cc := &CreditCard{balance: 500.00}
	btc := &Bitcoin{balance: 2.00}

	ProccessPayment(cc, 200.00)
	ProccessPayment(btc, 1.00)
}
