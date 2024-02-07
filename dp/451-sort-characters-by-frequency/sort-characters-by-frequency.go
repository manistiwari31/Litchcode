type pair struct {
	char  byte
	count int
}

func frequencySort(s string) string {
	count := make(map[byte]int)

	// Count the frequency of each character
	for i := 0; i < len(s); i++ {
		count[s[i]]++
	}

	// Create a slice of pairs from the map
	var pairs []pair
	for char, freq := range count {
		pairs = append(pairs, pair{char, freq})
	}

	// Sort the pairs based on their count in descending order
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].count > pairs[j].count
	})

	// Construct the sorted string
	var sortedString []byte
	for _, p := range pairs {
		for i := 0; i < p.count; i++ {
			sortedString = append(sortedString, p.char)
		}
	}

	return string(sortedString)
}
