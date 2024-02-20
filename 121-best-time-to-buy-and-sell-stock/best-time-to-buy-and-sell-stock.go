func maxProfit(prices []int) int {
    res := 0

    lowest := prices[0]

    for _,p := range prices {
        if p<lowest {
            lowest = p
        }
        if res<p-lowest {
            res = p-lowest
        }
    }
    return res
}