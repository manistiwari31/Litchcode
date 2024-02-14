func rearrangeArray(nums []int) []int {
        results := make([]int, len(nums))

        posIdx, negIdx := 0, 1
        for _, num := range nums {
                if num > 0 {
                        results[posIdx] = num
                        posIdx += 2
                } else {
                        results[negIdx] = num
                        negIdx += 2
                }
        }

        return results
}