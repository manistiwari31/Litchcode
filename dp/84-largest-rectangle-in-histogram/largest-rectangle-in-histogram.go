func largestRectangleArea(a []int) (res int) {
	a = append(a, -1)
	n := len(a)
	l := make([]int, n)
    for i := 0; i < n; i++ {
		j := i - 1
		for ; j > -1 && a[j] > a[i]; j = l[j] {
			res = max(a[j]*(i-l[j]-1), res)
		}
		l[i] = j
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}