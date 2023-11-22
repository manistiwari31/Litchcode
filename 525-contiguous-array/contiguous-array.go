func findMaxLength(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    for i := range nums {
        if nums[i] == 0 {
            nums[i] = -1
        }
    }

    mp := make(map[int]int)
    sum := 0
    ans := 0
    for i:=0; i<len(nums); i++ {
        sum += nums[i]
        if sum == 0 {
            ans = max(ans, i+1)
        } else if _, ok := mp[sum]; ok {
            ans = max(ans, i-mp[sum])
        } else {
            mp[sum] = i
        }
    }
    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
