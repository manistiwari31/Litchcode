func sortColors(nums []int) {
    for i := 0; i < len(nums)-1; i++ {
        for j := i; j < len(nums); j++ {
            if nums[i] > nums[j] {
                nums[i], nums[j] = nums[j], nums[i]
            }
        }
    }
}