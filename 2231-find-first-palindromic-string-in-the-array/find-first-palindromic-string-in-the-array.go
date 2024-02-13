func firstPalindrome(words []string) string {
    for _,s := range words {
        if Pali(s){
            return s
        }
        
    }
    return ""
}

func Pali(w string) bool {
    a, b := 0, len(w)-1;

    for (a<b) {
        if w[a]!=w[b] {
            return false
        }
        a++
        b--
    }
    return true
}