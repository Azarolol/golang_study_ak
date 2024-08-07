package main

import "fmt"

func mergeMaps(map1, map2 map[string]int) map[string]int {
	output := make(map[string]int)
	for k, v := range map1 {
		output[k] = v
	}
	for k, v := range map2 {
		output[k] = v
	}
	return output
}

func main() {
	map1 := map[string]int{"apple": 3, "banana": 2}
	map2 := map[string]int{"orange": 5, "grape": 4}
	mergedMap := mergeMaps(map1, map2)
	for key, value := range mergedMap {
		fmt.Printf("%s: %d\n", key, value)
	}
}
