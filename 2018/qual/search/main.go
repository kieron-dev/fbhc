package main

import (
	"fmt"
	"log"
)

func main() {
	var numCases int
	_, err := fmt.Scanf("%d", &numCases)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < numCases; i++ {
		var needle string
		_, err := fmt.Scanf("%s", &needle)
		if err != nil {
			log.Fatal(err)
		}
		counter := MakeCounterExample(needle)
		if Contains(needle, counter) {
			panic("oops")
		}
		fmt.Printf("Case #%d: %s\n", i+1, counter)
	}
}

func MakeCounterExample(needle string) string {
	for i := 1; i < len(needle); i++ {
		if needle[i] == needle[0] {
			for j := i; j < 2*i && j < len(needle); j++ {
				if needle[j] != needle[j-i] {
					return needle[:j] + needle[1:]
				}
			}
		}
	}
	return "Impossible"
}

func Contains(needle, haystack string) bool {
	var i, j int
	for {
		if i > len(needle)-1 {
			return true
		}
		if j > len(haystack)-1 {
			return false
		}
		if needle[i] == haystack[j] {
			i++
			j++
		} else if i == 0 {
			j++
		} else {
			i = 0
		}
	}
}
