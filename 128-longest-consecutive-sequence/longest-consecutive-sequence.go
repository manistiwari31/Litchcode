func longestConsecutive(nums []int) int {
	set := map[int]bool{}
	for _, num := range nums {
		set[num] = true
	}

	res := 0
	for _, num := range nums {
		if set[num-1] {
			continue
		}

		sequence, temp := 1, num + 1
		for set[temp] {
			sequence++
			temp++
		}

		res = max(res, sequence)
        if sequence > len(nums) / 2{
            break
        }
	}
	return res
}