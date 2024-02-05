func topKFrequent(nums []int, k int) []int {
	f := map[int]int{} // freq of each element
    maxFreq := 0
	for i := range nums {
		f[nums[i]]++
        maxFreq = max(maxFreq,f[nums[i]])
	}    

    buckets := make([][]int,maxFreq+1)
    for key,freq := range f {
        buckets[freq] = append(buckets[freq],key)
    }

    nums = make([]int,0)
    for i := maxFreq; i >= 0 && k > 0  ; i-- {
        k-= len(buckets[i]) // answer is unique; so it cannot be that we take 2 elements from same freq and leave third
        nums = append(nums,buckets[i]...)
    }
    return nums
}