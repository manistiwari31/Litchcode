func sortColors(nums []int) {
    lastZero := 0
    lastOne := 0

    for tracker := range nums {
        if nums[tracker] == 0 {
            nums[lastZero], nums[tracker] = nums[tracker], nums[lastZero]
            if lastZero < lastOne {
                nums[lastOne], nums[tracker] = nums[tracker], nums[lastOne]
            }
            lastZero++
            lastOne++
        } else if nums[tracker] == 1 {
            nums[lastOne], nums[tracker] = nums[tracker], nums[lastOne]
            lastOne++
        }
    }
}
