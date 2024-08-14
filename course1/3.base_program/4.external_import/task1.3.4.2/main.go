package main

import (
	"fmt"

	"github.com/ksrof/gocolors"
)

func ColorizeRed(a string) string {
	return gocolors.Red(a, "")
}

func ColorizeGreen(a string) string {
	return gocolors.Green(a, "")
}

func ColorizeBlue(a string) string {
	return gocolors.Blue(a, "")
}

func ColorizeYellow(a string) string {
	return gocolors.Yellow(a, "")
}

func ColorizeMagenta(a string) string {
	return gocolors.Magenta(a, "")
}

func ColorizeCyan(a string) string {
	return gocolors.Cyan(a, "")
}

func ColorizeWhite(a string) string {
	return gocolors.White(a, "")
}

func ColorizeCustom(a string, r, g, b uint8) string {
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm%s\x1b[0m", r, g, b, a)
}

func main() {
	fmt.Println(ColorizeRed("Red"))
	fmt.Println(ColorizeGreen("Green"))
	fmt.Println(ColorizeBlue("Blue"))
	fmt.Println(ColorizeYellow("Yellow"))
	fmt.Println(ColorizeMagenta("Magenta"))
	fmt.Println(ColorizeCyan("Cyan"))
	fmt.Println(ColorizeWhite("White"))
	fmt.Println(ColorizeCustom("Pink", 255, 192, 203))
}
