func largestOddNumber(num string) string {
    index := -1
    for i,u := range num {
        e := string(u)
        if e == "1" || e == "3" || e == "5" || e == "7" || e == "9" {
            index = i
        }
    }
    return num[:index+1]
}