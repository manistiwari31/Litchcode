func largestDivisibleSubset(nums []int) []int {
    res, subsets := []int{}, make([][]int, len(nums))

    sort.Ints(nums)

    for i := 0; i < len(nums); i++ {
        cur := []int{}

        for j := 0; j < i; j++ {
            if nums[i] % nums[j] == 0 {
                if len(subsets[j]) > len(cur) {
                    cur = subsets[j]
                }
            }
        }

        newSubset := make([]int, len(cur) + 1)
        copy(newSubset, cur)
        newSubset[len(cur)] = nums[i]
        subsets[i] = newSubset
        
        if len(subsets[i]) > len(res) {
             res = subsets[i]
        }
    }

    return res
}