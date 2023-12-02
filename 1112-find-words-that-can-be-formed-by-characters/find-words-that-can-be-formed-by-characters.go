func countCharacters(words []string, chars string) int {
    counter := make(map[rune]int)
    for _,letter := range(chars){
        counter[letter]++
    }
    answer := 0

    for _,word := range(words){
        word_counter := make(map[rune]int)
        for _,letter := range(word){
            word_counter[letter]++
        }
        good := true
        for key,value := range(word_counter){
            if _,ok := counter[key]; ok && value <= counter[key]{
                continue
            }
            good = false
            break
        }
        if good{
            answer+=len(word)
        }
    }
    return answer
}
