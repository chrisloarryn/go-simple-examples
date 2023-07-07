package main

import (
	"fmt"
	"strings"
)

/*
 * Complete the 'makeAnagram' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. STRING a
 *  2. STRING b
 */

func makeAnagram(a string, b string) int32 {
    a = strings.ToLower(a)
    b = strings.ToLower(b)

    freqA := make(map[rune]int)
    freqB := make(map[rune]int)

    for _, char := range a {
        freqA[char]++
    }

    for _, char := range b {
        freqB[char]++
    }

    deletions := int32(0)

    for char, countA := range freqA {
        countB := freqB[char]
        if countA != countB {
            deletions += intAbs(countA - countB)
        }
    }

    // Recorrer los caracteres que están presentes solo en la cadena b
    for char, countB := range freqB {
        countA := freqA[char]
        if countA == 0 {
            deletions += int32(countB)
        }
    }

    return deletions
}

// Función auxiliar para obtener el valor absoluto de un número entero
func intAbs(num int) int32 {
    if num < 0 {
        return -int32(num)
    }
    return int32(num)
}


func main() {
	// Ejemplo de uso
	cadenaA := "hola"
	cadenaB := "aloha"
	resultado := makeAnagram(cadenaA, cadenaB)
	fmt.Println(resultado) // Salida esperada: 0
}
