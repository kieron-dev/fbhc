package main

import (
	"fmt"
	"log"
	"sort"
	"strings"
)

func main() {
	var numCases int
	fmt.Scanf("%d\n", &numCases)

	for c := 0; c < numCases; c++ {
		var n, k, v int
		_, err := fmt.Scanf("%d %d %d\n", &n, &k, &v)
		if err != nil {
			log.Fatal(err)
		}

		attractions := make([]string, n)
		for i := 0; i < n; i++ {
			fmt.Scanf("%s\n", &attractions[i])
		}

		soln := getAttractions(attractions, n, k, v)
		fmt.Printf("Case #%d: %s\n", c+1, soln)
	}
}

func getAttractions(attractions []string, n, k, v int) string {
	start := ((v - 1) * k) % n
	if start < 0 {
		start += n
	}
	nums := []int{}
	for i := 0; i < k; i++ {
		nums = append(nums, (start+i)%n)
	}
	sort.Ints(nums)
	attrs := []string{}
	for _, i := range nums {
		attrs = append(attrs, attractions[i])
	}
	return strings.Join(attrs, " ")
}
