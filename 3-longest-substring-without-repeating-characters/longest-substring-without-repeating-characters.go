func lengthOfLongestSubstring(s string) int {
    view := make(map[byte]int)
    left := 0
    res := 0

    for right := range s  {
     	if _, found := view[s[right]]; found {
			for s[right] != s[left] {
				delete(view, s[left])
				left++
			} 
            left++
        } else {
            view[s[right]] = right
			if right-left+1 > res {
				res = right - left + 1
			}
        }  
    }
    return res
}