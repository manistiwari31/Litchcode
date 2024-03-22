
func sortColors(nums []int) {
    quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, low, high int) {
    if low < high {
        pivotIndex := partition(nums, low, high)
        quickSort(nums, low, pivotIndex-1)
        quickSort(nums, pivotIndex+1, high)
    }
}

func partition(nums []int, low, high int) int {
    pivot := nums[high]
    i := low - 1
    for j := low; j < high; j++ {
        if nums[j] < pivot {
            i++
            nums[i], nums[j] = nums[j], nums[i]
        }
    }
    nums[i+1], nums[high] = nums[high], nums[i+1]
    return i + 1
}
