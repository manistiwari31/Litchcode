func firstUniqChar(s string) int {
	freq := make([]int, 26)

	// Count occurrences of each character
	for _, ch := range s {
		freq[ch-'a']++
	}

	// Find the first unique character
	for i, ch := range s {
		if freq[ch-'a'] == 1 {
			return i
		}
	}

	return -1
}