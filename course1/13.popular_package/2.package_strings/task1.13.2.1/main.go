package main

import (
	"fmt"
	"strings"
)

func CountWordsInText(txt string, words []string) map[string]int {
	dict := make(map[string]int)
	for _, word := range words {
		dict[word] = 0
	}
	for _, word := range strings.Fields(txt) {
		_, ok := dict[strings.ToLower(word)]
		if ok {
			dict[strings.ToLower(word)]++
		}
	}
	return dict
}

func main() {
	txt := `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec a diam lectus. Sed sit amet ipsum mauris.
		Maecenas congue ligula ac quam viverra nec consectetur ante hendrerit. Donec et mollis dolor.
		Praesent et diam eget libero egestas mattis sit amet vitae augue.`
	words := []string{"sit", "amet", "lorem"}
	
	result := CountWordsInText(txt, words)

	fmt.Println(result) // map[amet: 2 lorem: 1 sit:3]
}