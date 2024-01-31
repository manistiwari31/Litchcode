func dailyTemperatures(temp []int) []int {
	stack := make([]int, 0)
	arr := make([]int, len(temp))

	for i := 0; i < len(temp); i++ {
		for len(stack) > 0 && temp[stack[len(stack)-1]] < temp[i] {
			ind := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			arr[ind] = i - ind
		}
		stack = append(stack, i)
	}

	return arr
}