
func evalRPN(tokens []string) int {
	stack := []int{}

	// Define anonymous functions for each operator
	operators := map[string]func(int, int) int{
		"*": func(a, b int) int { return a * b },
		"+": func(a, b int) int { return a + b },
		"-": func(a, b int) int { return a - b },
		"/": func(a, b int) int { return a / b },
	}

	for _, token := range tokens {
		if isOperator(token) {
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			// Use the anonymous function based on the operator
			operation := operators[token]
			stack = append(stack, operation(a, b))
		} else {
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}

	return stack[0]
}

func isOperator(token string) bool {
	return token == "*" || token == "+" || token == "-" || token == "/"
}
