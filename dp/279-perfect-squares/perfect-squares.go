func min(a, b int) int {
    if a > b {return b}
    return a
}

func numSquares(n int) (res int) {
    power, dp, j := make([]int, 101), make([]int, n+1), 0
    for i := 0; i < 100; i++ {power[i] = (i+1) * (i+1)}

    for i := 1; i <= n; i++ {
        dp[i], j = math.MaxInt, 0
        for power[j] <= i {dp[i], j = min(dp[i], dp[i-power[j]]+1), j+1}
    }
    return dp[n]
}