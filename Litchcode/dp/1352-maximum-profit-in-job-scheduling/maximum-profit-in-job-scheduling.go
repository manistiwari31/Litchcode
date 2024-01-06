// Slices implements the Go standard lib's Sort interface
// by having Len, Less, and Swap methods.
type Slices [][]int

func (s Slices) Len() int { return len(s[0]) }
func (s Slices) Less(i, j int) bool { return s[0][i] < s[0][j] }
func (s Slices) Swap(i, j int) {
	for sub := range s {
		s[sub][i], s[sub][j] = s[sub][j], s[sub][i]
	}
}

func jobScheduling(startTime, endTime, profit []int) int {
    sort.Sort(Slices{endTime, startTime, profit})

    dp := make([]int, len(startTime))
    dp[0] = profit[0]

    for i := 1; i < len(startTime); i++ {
        jobProfit := profit[i]
        firstOverlappingJobIdx := sort.Search(i, func(j int) bool {
            return startTime[i] < endTime[j]
        })
        if firstOverlappingJobIdx > 0 {
            jobProfit += dp[firstOverlappingJobIdx-1]
        }

        // dp[i-1] = max profit if do *not* take this job
        dp[i] = max(dp[i-1], jobProfit)
    }
    return dp[len(dp)-1]
}

func max(a, b int) int {
    if a > b { return a }
    return b
}