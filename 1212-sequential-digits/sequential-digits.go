func sequentialDigits(low int, high int) []int {
	res := make([]int, 0)
	val := initialize(low)
	for {
		if val > high {
			break
		}
		if val >= low {
			res = append(res, val)
		}
		val = next(val)
	}
	return res

}

func initialize(low int) int {
	digits := getDigits(low)
	return makeByN(low/int(math.Pow10(digits-1)), digits)
}

func next(cur int) int {
	if cur%10 == 9 {
		return makeByN(1, getDigits(cur)+1)
	}
	var res int
	var cnt int
	for cur > 0 {
		res += (1 + cur%10) * int(math.Pow10(cnt))
		cnt++
		cur /= 10
	}
	return res
}

func makeByN(start, n int) int {
	if n > 9 {
		return math.MaxInt
	}
	if start+n > 10 {
		return makeByN(1, n+1)
	}
	var res int
	for i := 1; i <= n; i++ {
		res = res*10 + start
		start++
	}
	return res
}

func getDigits(val int) int {
	var cnt int
	for val > 0 {
		val /= 10
		cnt++
	}
	return cnt
}
