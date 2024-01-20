func decodeString(s string) string {
	stack := []rune{} 

	for _, ch :=range s{
		if ch != ']' {
			stack = append(stack, ch)
		} else {
			subStr := ""
			for stack[len(stack)-1] != '[' {
				subStr = string(stack[len(stack)-1]) + subStr
				stack = stack[:len(stack)-1] // pop
			}
		
        
			stack = stack[:len(stack)-1] // pop'['

			numStr := ""
			for len(stack) > 0 && '0' <= stack[len(stack)-1] && stack[len(stack)-1] <='9' {
				numStr =  string(stack[len(stack)-1]) + numStr
				stack = stack[:len(stack)-1] // pop
			}
			
			num, _ := strconv.Atoi(numStr)
			for num > 0 {
				stack = append(stack, []rune(subStr)...)
				num--
			}
		}
	}

	return string(stack)
}