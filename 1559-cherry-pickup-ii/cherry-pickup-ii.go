func cherryPickup(grid [][]int) int {
    row := len(grid)
    col := len(grid[0])

    dp := make([][][]int, row)

    for i := 0; i < row; i++ {
        dp[i] = make([][]int, col)
        for j := 0; j < col; j++ {
            dp[i][j] = make([]int, col)
            for k := 0; k < col; k++ {
                dp[i][j][k] = -1
            }
        }
    }

    return solve(0, 0, col-1, grid, dp)
}

func solve(i int, j int, k int, grid [][] int, dp [][][] int) int {
    if i == len(grid) {
        return 0    //We are outside the grid
    }

    if j < 0 || k < 0 || j >= len(grid[0]) || k >= len(grid[0]) {
        return 0    //We are outside the grid
    }

    if dp[i][j][k] != -1 {
        return dp[i][j][k]  //We have already calculated the subproblem, lets just return it
    }

    //fmt.Printf("%d %d %d\n", i, j, k)

    directions := []int{-1,0,1}

    max := 0
    for x := 0; x < 3; x++ {
        for y := 0; y < 3; y++ {
            max = Max(max, solve(i+1, j+directions[x], k+directions[y], grid, dp))
        }
    }

    soFar := grid[i][j] + grid[i][k]    //What we collect from the current row

    if j == k {
        soFar = grid[i][j]  //Both robots landed on the same cell, so we would count cherries only once
    }

    dp[i][j][k] = soFar + max
    return dp[i][j][k]
}

func Max(a int, b int) int {
    return int(math.Max(float64(a), float64(b)))
}