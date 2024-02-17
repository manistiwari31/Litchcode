func rotate(nums []int, k int)  {
   l := len(nums)
   k = k%l

   rev(nums, 0, l-1)
   rev(nums,0,k-1)
   rev(nums,k,l-1)

}

func rev (nums[]int, st, en int) {
    for st<en {
        nums[st], nums[en] = nums[en], nums[st]
        st++
        en--
    }
}
