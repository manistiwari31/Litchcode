func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0, len(asteroids))
	last := -1

    for true {
        if last >= 1 && stack[last] < 0 && stack[last] * stack[last-1] < 0 {
            if abs(stack[last]) == abs(stack[last-1]) {
                stack = stack[:last-1]
                last -= 2
            } else {
                stack[last-1] = max(stack[last], stack[last-1])
                stack = stack[:last]
                last--
            }
        } else {
            if len(asteroids) == 0 {
                break
            } else {
                stack = append(stack, asteroids[0])
                asteroids = asteroids[1:]
                last++
            }
        }
    }

	return stack
}

func max(a, b int) int {
	if abs(a) > abs(b) {
		return a
	}

	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}