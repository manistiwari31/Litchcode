func topKFrequent(nums []int, k int) []int {
	valueMap := make(map[int][]int)
	hashMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if tt, vv := hashMap[nums[i]]; vv {
			hashMap[nums[i]] = tt + 1
			continue
		}
		hashMap[nums[i]] = 1
	}
	for kk, v := range hashMap {
		valueMap[v] = append(valueMap[v], []int{kk}...)
	}
	valueList := make([]int, 0)
	result := make([]int, 0)
	count := k
	keys := make([]int, 0)
	for k, _ := range valueMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, v := range keys {
		valueList = append(valueList, valueMap[v]...)
	}
	for i := len(valueList) - 1; i >= 0; i-- {
		if count > 0 {
			result = append(result, valueList[i])
		}
		count--
	}
	return result
}