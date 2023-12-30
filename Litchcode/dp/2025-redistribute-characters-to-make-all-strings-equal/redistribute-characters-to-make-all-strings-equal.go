func makeEqual(words []string) bool {
    counts := make(map[rune]int)
    for _,word:= range(words){
        for _,c := range(word){
            counts[c]++
        }
    }

    for _, count := range counts {
        if count%len(words) != 0 {
            return false
        }
    }
    return true
}