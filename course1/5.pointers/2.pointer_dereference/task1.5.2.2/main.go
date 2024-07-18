package main

import "fmt"

func Factorial(n *int) int {
	ans := 1
	for i := 2; i <= *n; i++ {
		ans *= i
	}
	return ans
}

func isPalindrome(str *string) bool {
	runes := []rune(*str)
	for i := 0; i <= len(runes)/2; i++ {
		if runes[i] != runes[len(runes)-i-1] {
			return false
		}
	}
	return true
}

func CountOccurences(numbers *[]int, target *int) int {
	count := 0
	for _, num := range *numbers {
		if num == *target {
			count++
		}
	}
	return count
}

func ReverseString(str *string) string {
	newStr := ""
	for _, c := range *str {
		newStr = string(c) + newStr
	}
	return newStr
}

func main() {
	a := 6
	fmt.Println(Factorial(&a))
	b := "tenet"
	fmt.Println(isPalindrome(&b))
	c := []int{0, 1, 2, 3, 2}
	d := 2
	fmt.Println(CountOccurences(&c, &d))
	e := "string to be reversed"
	fmt.Println(ReverseString(&e))
}
