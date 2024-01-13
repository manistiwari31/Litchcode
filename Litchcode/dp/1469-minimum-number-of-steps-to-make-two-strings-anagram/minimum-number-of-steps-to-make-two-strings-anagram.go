func minSteps(s string, t string) int {
    counts := make(map[rune]int)
    for _,c := range(s){
        counts[c]++
    }
    countt := make(map[rune]int)
    for _,c := range(t){
        countt[c]++
    }
    res := len(s)
    for char, count := range counts {
        if countt[char]>0 {
            if count>countt[char]{
                res -= countt[char]
            } else {
                res-=count
            }
        }
    }
    return res    
}