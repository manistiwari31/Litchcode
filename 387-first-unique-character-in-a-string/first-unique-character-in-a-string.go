func firstUniqChar(s string) int {
    countMap := map[rune]int{}

    // Count occurrences of each character
    for _, t := range s {
        countMap[t]++
    }

    // Find the first unique character
    for i, t := range s {
        if countMap[t] == 1 {
            return i
        }
    }

    return -1
}
