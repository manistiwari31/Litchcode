func plusOne(digits []int) []int {
    if len(digits) == 0 {
        return []int{1}
    }
    
    for o := len(digits)-1; o >= 0; o-- {
        if digits[o] == 9 {
            digits[o] = 0
            if o == 0 {
                digits = append([]int{1},digits...)
            }
        } else {
            digits[o] += 1
            return digits
        }
    }
    return digits
}
