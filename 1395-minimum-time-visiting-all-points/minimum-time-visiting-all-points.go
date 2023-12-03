func minTimeToVisitAllPoints(points [][]int) int {
    if len(points) < 2 { return 0 }
    result := 0

    for i := 1; i < len(points); i++  {
        result += timeToTravel(points[i], points[i-1])
    }

    return result
}

func timeToTravel(point1, point2 []int) int {
    xdiff := abs(point2[0] - point1[0])
    ydiff := abs(point2[1] - point1[1])
    return max(xdiff, ydiff)
}

func abs(x int) int {
    if x < 0 { return -x }
    return x
}

func max(x, y int) int {
    if x > y { return x }
    return y
}