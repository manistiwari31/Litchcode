func topKFrequent(nums []int, k int) []int {
    m := make(map[int]int)
    for _, val := range nums {
        m[val] = m[val]+1
    }
    pq := make([]Item, k)
    i := 0
    for key, val := range m {
        if i<len(pq) {
            pq[i] = Item{value: key, count: val}
            i++
        } else {
            min := math.MaxInt32
            minInd := -1
            for j, v := range pq {
                if v.count < min {
                    min = v.count
                    minInd = j
                }
            }
            if val > min {
                pq[minInd] = Item{value: key, count: val}
            }
        }
    }
    sort.Slice(pq, func(i, j int) bool { return pq[i].count > pq[j].count })
    res := make([]int, k)
    for j, v := range pq {
        res[j] = v.value
    }            
    return res
}

type Item struct {
	value int 
	count int   
}