package main

import (
	"fmt"
	"unsafe"
)

func sizeOfBool(b bool) int {
	return int(unsafe.Sizeof(b))
}

func sizeOfInt(n int) int {
	return int(unsafe.Sizeof(n))
}

func sizeOfInt8(n int8) int {
	return int(unsafe.Sizeof(n))
}

func sizeOfInt16(n int16) int {
	return int(unsafe.Sizeof(n))
}

func sizeOfInt32(n int32) int {
	return int(unsafe.Sizeof(n))
}

func sizeOfInt64(n int64) int {
	return int(unsafe.Sizeof(n))
}

func sizeOfUint(n uint) int {
	return int(unsafe.Sizeof(n))
}

func sizeOfUint8(n uint8) int {
	return int(unsafe.Sizeof(n))
}

func main() {
	fmt.Printf("bool size: %d\n", sizeOfBool(true))
	fmt.Printf("int size: %d\n", sizeOfInt(1))
	fmt.Printf("int8 size: %d\n", sizeOfInt8(1))
	fmt.Printf("int16 size: %d\n", sizeOfInt16(1))
	fmt.Printf("int32 size: %d\n", sizeOfInt32(1))
	fmt.Printf("int64 size: %d\n", sizeOfInt64(1))
	fmt.Printf("uint size: %d\n", sizeOfUint(1))
	fmt.Printf("uint8 size: %d\n", sizeOfUint8(1))
}
