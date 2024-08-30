package main

import (
	"container/list"
	"fmt"
)

type Car struct {
	LicensePlate string
}

type ParkingLot struct {
	space *list.List
}

func NewParkingLot() *ParkingLot {
	return &ParkingLot{space: list.New()}
}

func (p *ParkingLot) Park(c Car) {
	p.space.PushFront(c)
	fmt.Println("Автомобиль", c.LicensePlate, "припаркован.")
}

func (p *ParkingLot) Leave() {
	if p.space.Len() == 0 {
		fmt.Println("Парковка пуста.")
		return
	}
	c := p.space.Back()
	p.space.Remove(c)
	fmt.Println("Автомобиль", c.Value.(Car).LicensePlate, "покинул парковку.")
}

func main() {
	parkingLot := NewParkingLot()
	parkingLot.Park(Car{LicensePlate: "ABC-123"})
	parkingLot.Park(Car{LicensePlate: "XYZ-789"})
	parkingLot.Leave()
	parkingLot.Leave()
	parkingLot.Leave()
}
