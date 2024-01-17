func uniqueOccurrences(arr []int) bool {
    counts := make(map[rune]int)
   
        for _,c := range(arr){
            counts[rune(c)]++
        }
    
    type void struct{}
    var insert void
    set := make(map[string]void)

    for _, count := range counts {
        if _, ok := set[string(count)]; ok {
            return false
        } else {
            set[string(count)]= insert
            }
        }

    return true

}