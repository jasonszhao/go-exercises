// solution to https://tour.golang.org/moretypes/23
package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	result := make(map[string]int)
	
	for _, word := range words {
		result[word]++
	}
	return result
}

func main() {
	wc.Test(WordCount)
}

