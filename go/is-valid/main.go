package main

import "fmt"

type Options struct {
    Text string
    Result string
}

func main(){
	// map with key and value type string declared and initialized
    optionsMap := map[string]Options{
        "test invalid 1": {
            Text: "aabbcd",
            Result: "NO",
        },
        "test invalid 2": {
            Text: "aabbccddeefghi",
            Result: "NO",
        },
        "test valid 1": {
            Text: "abcdefghhgfedecba",
            Result: "YES",
        },
    }

    for key, value := range optionsMap {
        result := isValid(value.Text)
        if result != value.Result {
            //print and salt line
            fmt.Println(fmt.Sprintf("Test %s failed. Expected %s, got %s", key, value.Result, result))
            return
        } else {
            fmt.Println(fmt.Sprintf("Test %s passed. Expected %s, got %s", key, value.Result, result))
        }
    }
}

func isValid(s string) string {
    frecuencia := make(map[rune]int)

    for _, char := range s {
        frecuencia[char]++
    }

    frecuenciaCount := make(map[int]int)

    for _, freq := range frecuencia {
        frecuenciaCount[freq]++
    }

    frecuenciasUnicas := make([]int, 0, len(frecuenciaCount))
    for freq := range frecuenciaCount {
        frecuenciasUnicas = append(frecuenciasUnicas, freq)
    }

    if len(frecuenciasUnicas) > 2 {
        return "NO"
    }

    if len(frecuenciasUnicas) == 1 {
        return "YES"
    }

    freq1, freq2 := frecuenciasUnicas[0], frecuenciasUnicas[1]

    if (frecuenciaCount[freq1] == 1 && freq1 == 1) || (frecuenciaCount[freq2] == 1 && freq2 == 1) {
        return "YES"
    }

    if diff := freq1 - freq2; (diff == 1 || diff == -1) && (frecuenciaCount[freq1] == 1 || frecuenciaCount[freq2] == 1) {
        return "YES"
    }

    return "NO"
}
