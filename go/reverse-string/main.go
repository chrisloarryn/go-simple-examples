package main

import "fmt"

var (
	value = "this is a text that will be reversed"
)

func main() {
	reversed := reverse(value)
	fmt.Println(reversed)
}

func reverse(s string) string {
	var reversed string
	for i := len(s) - 1; i >= 0; i-- {
		reversed += string(s[i])
	}
	return reversed
}
