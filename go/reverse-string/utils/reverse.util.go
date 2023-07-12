package utils

import "fmt"

const (
	ErrEmpty = "empty string"
)

func Reverse(s string) (string, error) {
	if len(s) == 0 {
		return "", fmt.Errorf(ErrEmpty)
	}

	var reversed string
	for i := len(s) - 1; i >= 0; i-- {
		reversed += string(s[i])
	}
	return reversed, nil
}
