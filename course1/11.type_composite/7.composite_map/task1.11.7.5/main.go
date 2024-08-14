package main

import (
	"fmt"
	"strings"
)

func filterSentence(sentence string, filter map[string]bool) string {
	newSentence := strings.Builder{}
	for _, word := range strings.Split(sentence, " ") {
		_, ok := filter[word]
		if !ok {
			newSentence.WriteString(word + " ")
		}
	}
	return strings.TrimRight(newSentence.String(), " ")
}

func main() {
	sentence := "Lorem ipsum dolor sit amet consectetur adipiscing elit ipsum"
	filter := map[string]bool{"ipsum": true, "elit": true}

	filteredSentence := filterSentence(sentence, filter)
	fmt.Println(filteredSentence)
}
