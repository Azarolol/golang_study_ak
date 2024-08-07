package main

import (
	"fmt"
	"strings"
)

func countWordOccurences(text string) map[string]int {
	output := make(map[string]int)
	for _, word := range strings.Split(text, " ") {
		_, ok := output[word]
		if ok {
			output[word]++
		} else {
			output[word] = 1
		}
	}
	return output
}

func main() {
	text := "Lorem ipsum dolor sit amet consectetur adipiscing elit ipsum"
	occurences := countWordOccurences(text)

	for word, count := range occurences {
		fmt.Printf("%s: %d\n", word, count)
	}
}
