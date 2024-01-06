func maxInt(i, j int) int {
	if i > j {
		return i
	}
	return j
}
func dpSolution(nums []int) int {
	n := len(nums)
	// the longest length from 0 to i
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if nums[i] < nums[j] {
				dp[j] = maxInt(dp[i]+1, dp[j])
			}
		}
	}
	max := 0
	for i := 0; i < n; i++ {
		max = maxInt(max, dp[i])
	}
	return max
}
// refer to https://en.wikipedia.org/wiki/Patience_sorting
func patienceSortingSolution(nums []int) int {
	n := len(nums)
	top := make([][]int, 0)
	for i := 0; i < n; i++ {
		// binary search
		l, r := 0, len(top)-1
		for l <= r {
			mid := l + (r-l)/2
			numOfTop := top[mid][len(top[mid])-1]
			if nums[i] <= numOfTop {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
		if l >= len(top) {
			top = append(top, []int{nums[i]})
			continue
		}
		top[l] = append(top[l], nums[i])
	}
	return len(top)
}
func lengthOfLIS(nums []int) int {
	return patienceSortingSolution(nums)
}