package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func binaryStringToFloat(binary string) float32 {
	var number uint32
	// Преобразование строки в двоичной системе в целочисленное представление
	number64, _ := strconv.ParseUint(binary, 2, 32)
	number = uint32(number64)
	// Преобразование целочисленного представления в число с плавающей точкой
	floatNumber := *(*float32)(unsafe.Pointer(&number))
	return floatNumber
}

func main() {
	fmt.Println(binaryStringToFloat("00111110001000000000000000000000")) // 0.15625
}
