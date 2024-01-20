func minAddToMakeValid(s string) int {
    stack := []rune{}

	for _, ch := range s {
		 if ch == ')' && len(stack) > 0 && stack[len(stack)-1] == '(' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, ch)
		}
	}

	return len(stack)
}