func twoSum(nums []int, target int) []int {
    magic := map[int]int{}
    for i1, v1 := range nums {
        if i2, ok := magic[target-v1]; ok {
            return []int{i2,i1}
        }
        magic[v1] = i1
    }
    return nil
}