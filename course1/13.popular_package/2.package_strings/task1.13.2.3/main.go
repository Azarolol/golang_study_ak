package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func GenerateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	output := strings.Builder{}
	for i := 0; i < length; i++ {
		r := rand.Int31n(1000)
		output.WriteRune(r)
	}
	return output.String()
}

func main() {
	randomString := GenerateRandomString(10)
	fmt.Println(randomString) // Выводит случайную строку длиной 10 символов, например "3aBcD9eFgH"
}
