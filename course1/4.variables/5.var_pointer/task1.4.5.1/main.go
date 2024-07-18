package main

import "fmt"

func changeInt(a *int) {
	*a = 20
}

func changeFloat(b *float64) {
	*b = 6.28
}

func changeString(c *string) {
	*c = "Goodbye, world!"
}

func changeBool(d *bool) {
	*d = false
}

func main() {
	a := 19
	b := 6.27
	c := "Hello, world!"
	d := true
	changeInt(&a)
	fmt.Println(a)
	changeFloat(&b)
	fmt.Println(b)
	changeString(&c)
	fmt.Println(c)
	changeBool(&d)
	fmt.Println(d)
}
