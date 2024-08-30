package main

import (
	"fmt"
	"sort"
)

func sortTheStudents(score [][]int, k int) [][]int {
	sort.Slice(score, func(i, j int) bool {
		return score[i][k] > score[j][k]
	})
	return score
}

func main() {
	score := [][]int{{10, 6, 9, 1}, {7, 5, 11, 2}, {4, 8, 3, 15}}
	k := 2
	sorted := sortTheStudents(score, k)
	fmt.Println(sorted)
}
