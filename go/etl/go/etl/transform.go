package etl

import "strings"

func Transform(crufty map[int][]string) map[string]int {
	result := make(map[string]int)
	for score, letters := range crufty {
		for _, letter := range letters {
			result[strings.ToLower(letter)] = score
		}
	}
	return result
}