package main

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	routeMap := make(map[int]struct{})
	for _, edge := range edges {
		routeMap[edge[1]] = struct{}{}
	}
	out := make([]int, 0)
	for i := 0; i < n; i++ {
		if _, ok := routeMap[i]; !ok {
			out = append(out, i)
		}
	}

	return out
}

var _ = findSmallestSetOfVertices(0, nil)
