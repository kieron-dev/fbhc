package main

import (
	"fmt"
	"log"
)

func main() {
	var numCases int
	_, err := fmt.Scanf("%d\n", &numCases)
	if err != nil {
		log.Fatal(err)
	}

	for c := 0; c < numCases; c++ {
		var degree int
		var coeffs []int
		_, err := fmt.Scanf("%d\n", &degree)
		if err != nil {
			log.Fatal(err)
		}

		for i := 0; i < degree+1; i++ {
			var coeff int
			_, err := fmt.Scanf("%d\n", &coeff)
			if err != nil {
				log.Fatal(err)
			}
			coeffs = append(coeffs, coeff)
		}
		soln := solve(coeffs)
		fmt.Printf("Case #%d: %+v\n", c+1, len(soln))
		for _, r := range soln {
			fmt.Printf("%f\n", r)
		}
	}
}

func solve(coeffs []int) []float64 {
	// a * x ^ 3 + b * x ^ 2 + c * x ^ 1 + d * x ^ 0 =	(a * x) ^ (((3 + b) * x) ^ ((2 + c) * x))
	fmt.Printf("coeffs = %+v\n", coeffs)
	return []float64{0.0}
}
