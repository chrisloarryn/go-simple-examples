package utilities

func SubstrCount(n int32, s string) int64 {
	count := int64(n) // Each individual character is a valid substring

	// Count special substrings of the form "aaa...a" where all characters are the same
	for i := 0; i < len(s); i++ {
		j := i + 1
		for j < len(s) && s[i] == s[j] {
			count++
			j++
		}
	}

	// Count special substrings of the form "aba", where the middle character is the same
	for i := 1; i < len(s)-1; i++ {
		if s[i-1] == s[i+1] && s[i] != s[i-1] {
			k := 1
			for i-k >= 0 && i+k < len(s) && s[i-1] == s[i-k] && s[i+1] == s[i+k] {
				count++
				k++
			}
		}
	}

	return count
}
