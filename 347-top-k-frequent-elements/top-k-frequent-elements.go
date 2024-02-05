
func topKFrequent(nums []int, k int) []int {
    if k == 0 {
        return []int{}
    }
    countMap := make(map[int]int)

    for _,n := range nums {
        countMap[n] += 1
    }

    frequencyMap := make(map[int][]int)
    for key,v := range countMap {
        frequencyMap[v] = append(frequencyMap[v], key)
    }

    result := []int{}
    for i := len(nums); i > 0; i-- {
        if len(result) >= k {
            break
        }

        result = append(result, frequencyMap[i]...)       
    }

    return result
}