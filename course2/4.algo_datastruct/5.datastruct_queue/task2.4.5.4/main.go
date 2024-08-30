package main

import "fmt"

var stack []int

func push(value int) {
	stack = append(stack, value)
}

func pop() int {
	out := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	return out
}

func main() {
	// Пример использования стека для операций
	push(5)
	push(3)
	result := pop() + pop()
	push(result)
	// Вывод результата выполнения программы
	fmt.Println(stack[0]) // 8
}
