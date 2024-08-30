package main

import "fmt"

func numTilePossibilities(tiles string) int {
	dict := make(map[rune]int)
	for _, tile := range tiles {
		dict[tile]++
	}
	return walk(dict)
}

func walk(dict map[rune]int) int {
	out := 0
	for key := range dict {
		newDict := make(map[rune]int)
		for k, v := range dict {
			newDict[k] = v
		}
		out++
		newDict[key]--
		if newDict[key] == 0 {
			delete(newDict, key)
		}
		out += walk(newDict)
	}
	return out
}

func main() {
	tiles := "AAABBC"
	fmt.Println(numTilePossibilities(tiles))
}
