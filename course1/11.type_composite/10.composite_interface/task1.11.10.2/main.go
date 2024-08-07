package main

import (
	"fmt"
	"strings"
)

func Operate(f func(xs ...interface{}) interface{}, i ...interface{}) interface{} {
	return f(i...)
}

func Concat(xs ...interface{}) interface{} {
	sb := strings.Builder{}
	for _, x := range xs {
		sb.WriteString(x.(string))
	}
	return sb.String()
}

func Sum(xs ...interface{}) interface{} {
	if len(xs) == 0 {
		return interface{}(nil)
	}
	switch xs[0].(type) {
	case int:
		result := 0
		for _, x := range xs {
			result += x.(int)
		}
		return result
	case float64:
		result := 0.0
		for _, x := range xs {
			result += x.(float64)
		}
		return result
	default:
		return interface{}(nil)
	}
}

func main() {
	fmt.Println(Operate(Concat, "Hello, ", "World!"))  // Вывод: "Hello, World!"
	fmt.Println(Operate(Sum, 1, 2, 3, 4, 5))           // Вывод: 15
	fmt.Println(Operate(Sum, 1.1, 2.2, 3.3, 4.4, 5.5)) // Вывод: 16.5
}
