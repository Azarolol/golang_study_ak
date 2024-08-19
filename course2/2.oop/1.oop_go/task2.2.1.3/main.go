package main

import (
	"fmt"
)

type Order interface {
	AddItem(item string, quantity int) error
	RemoveItem(item string) error
	GetOrderDetails() map[string]int
}

type DineInOrder struct {
	orderDetails map[string]int
}

type TakeAwayOrder struct {
	orderDetails map[string]int
}

func (o *DineInOrder) AddItem(item string, quantity int) error {
	if quantity < 1 {
		return fmt.Errorf("can't add quantity less than 1 of item to order")
	}
	o.orderDetails[item] += quantity
	return nil
}

func (o *TakeAwayOrder) AddItem(item string, quantity int) error {
	if quantity < 1 {
		return fmt.Errorf("can't add quantity less than 1 of item to order")
	}
	o.orderDetails[item] += quantity
	return nil
}

func (o *DineInOrder) RemoveItem(item string) error {
	_, ok := o.orderDetails[item]
	if !ok {
		return fmt.Errorf("no such item in order")
	}
	delete(o.orderDetails, item)
	return nil
}

func (o *TakeAwayOrder) RemoveItem(item string) error {
	_, ok := o.orderDetails[item]
	if !ok {
		return fmt.Errorf("no such item in order")
	}
	delete(o.orderDetails, item)
	return nil
}

func (o *DineInOrder) GetOrderDetails() map[string]int {
	return o.orderDetails
}

func (o *TakeAwayOrder) GetOrderDetails() map[string]int {
	return o.orderDetails
}

func ManageOrder(o Order) {
	o.AddItem("Pizza", 2)
	o.AddItem("Burger", 1)
	o.RemoveItem("Pizza")
	fmt.Println(o.GetOrderDetails())
}

func main() {
	dineIn := &DineInOrder{orderDetails: make(map[string]int)}
	takeAway := &TakeAwayOrder{orderDetails: make(map[string]int)}

	ManageOrder(dineIn)
	ManageOrder(takeAway)
}
