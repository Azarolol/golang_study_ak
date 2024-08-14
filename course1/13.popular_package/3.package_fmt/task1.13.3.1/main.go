package main

import "fmt"

func generateMathString(operands []int, operator string) string {
	if len(operands) == 0 {
		return ""
	}
	answer := 0
	switch operator {
	case "+":
		for _, o := range operands {
			answer += o
		}
	case "-":
		answer = operands[0]
		for i := 1; i < len(operands); i++ {
			answer -= operands[i]
		}
	case "*":
		answer = 1
		for i := 0; i < len(operands); i++ {
			answer *= operands[i]
		}
	case "/":
		answer = operands[0]
		for i := 1; i < len(operands); i++ {
			answer /= operands[i]
		}
	default:
		return "Unknown operator"
	}
	output := ""
	for i := 0; i < len(operands)-1; i++ {
		output += fmt.Sprintf("%d %s ", operands[i], operator)
	}

	return fmt.Sprintf("%s%d = %d", output, operands[len(operands)-1], answer)
}

func main() {
	fmt.Println(generateMathString([]int{2, 4, 6}, "+")) // "2 + 4 + 6 = 12"
}
