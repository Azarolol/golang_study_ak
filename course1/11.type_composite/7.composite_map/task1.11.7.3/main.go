package main

import (
	"fmt"
	"strings"
)

func createUniqueText(text string) string {
	dict := make(map[string]struct{})
	newText := strings.Builder{}
	for _, word := range strings.Split(text, " ") {
		_, ok := dict[word]
		if !ok {
			newText.WriteString(word + " ")
			dict[word] = struct{}{}
		}
	}
	return strings.TrimRight(newText.String(), " ")
}

func main() {
	fmt.Println(createUniqueText("bar bar bar foo foo baz"))
}
