package utilities_test

import (
	"fmt"
	utilities "test-substring-count/utilities"
	"testing"
)

type Options struct {
	Number   int32
	Text     string
	Expected int64
}

const (
	Passed = "PASSED"
	Failed = "FAILED"
)

func TestSubstrCount(t *testing.T) {
	inputs := map[string]Options{
		"valid 1": {
			Number:   7,
			Text:     "abcbaba",
			Expected: 10,
		},
		"valid 2": {
			Number:   4,
			Text:     "aaaa",
			Expected: 10,
		},
		"valid 3": {
			Number:   5,
			Text:     "asasd",
			Expected: 7,
		},
	}
	for key, value := range inputs {
		result := utilities.SubstrCount(value.Number, value.Text)
		if result != value.Expected {
			//print and salt line
			fmt.Println(fmt.Sprintf("Test %s %s. Expected %d, got %d", key, Failed, value.Expected, result))
		} else {
			fmt.Println(fmt.Sprintf("Test %s %s. Expected %d, got %d", key, Passed, value.Expected, result))
		}
	}
}
