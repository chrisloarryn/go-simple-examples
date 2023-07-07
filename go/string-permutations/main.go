package main

import (
	"fmt"
	"reflect"
)

func substringsPermutations(A, B string) []string {
	var result []string
	lenA, lenB := len(A), len(B)
	counterB := make(map[rune]int)

	for _, char := range B {
		counterB[char]++
	}

	for i := 0; i <= lenA-lenB; i++ {
		substring := A[i : i+lenB]
		counterSub := make(map[rune]int)

		for _, char := range substring {
			counterSub[char]++
		}

		if reflect.DeepEqual(counterSub, counterB) {
			result = append(result, substring)
		}
	}

	return result
}

func main() {
	A := "cbbabcdbbb"
	B := "abbc"
	result := substringsPermutations(A, B)

	fmt.Printf("Result: %v\n", result)
}
