package main

import "fmt"

type TV interface {
	switchOn() bool
	switchOFF() bool
	GetStatus() bool
	GetModel() string
}

type LGTV struct {
	status bool
	model  string
}

type SamsungTV struct {
	status bool
	model  string
}

func (tv *SamsungTV) switchOn() bool {
	tv.status = true
	return tv.status
}

func (tv *LGTV) switchOn() bool {
	tv.status = true
	return tv.status
}

func (tv *SamsungTV) switchOFF() bool {
	tv.status = false
	return !tv.status
}

func (tv *LGTV) switchOFF() bool {
	tv.status = false
	return !tv.status
}

func (tv *SamsungTV) GetModel() string {
	return tv.model
}

func (tv *LGTV) GetModel() string {
	return tv.model
}

func (tv *SamsungTV) GetStatus() bool {
	return tv.status
}

func (tv *LGTV) GetStatus() bool {
	return tv.status
}

func (tv *SamsungTV) SamsungHub() string {
	return "SamsungHub"
}

func (tv *LGTV) LGHub() string {
	return "LGHub"
}

func main() {
	tv := &SamsungTV{
		status: true,
		model:  "Samsung XL-100500",
	}
	fmt.Println(tv.GetModel())   // Samsung XL-100500
	fmt.Println(tv.SamsungHub()) // SamsungHub
	fmt.Println(tv.GetStatus())  // true
	fmt.Println(tv.switchOFF())  // true
	fmt.Println(tv.GetStatus())  // false
	fmt.Println(tv.switchOn())   // true
	fmt.Println(tv.GetStatus())  // true
}
