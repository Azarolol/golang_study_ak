package main

import (
	"fmt"
	"sort"
)

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	out := make([]bool, 0, len(l))
	for i := 0; i < len(l); i++ {
		buf := make([]int, 0, r[i]+1-l[i])
		for j := l[i]; j <= r[i]; j++ {
			buf = append(buf, nums[j])
		}
		out = append(out, checkArray(buf))
	}
	return out
}

func checkArray(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	sort.Ints(nums)
	diff := nums[1] - nums[0]
	for i := 2; i < len(nums); i++ {
		if nums[i]-nums[i-1] != diff {
			return false
		}
	}
	return true
}

func main() {
	nums := []int{4, 6, 5, 9, 3, 7}
	l := []int{0, 0, 2}
	r := []int{2, 3, 5}
	fmt.Println(checkArithmeticSubarrays(nums, l, r))
}
