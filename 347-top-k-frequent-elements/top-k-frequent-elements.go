func topKFrequent(nums []int, k int) []int {
    var res []int
    
    frequencyMap:= make(map[int]int)
    
    for i := 0; i < len(nums); i++ {
        frequencyMap[nums[i]]++
    }
    
    var minHeap [][2]int
    
    for key, val := range frequencyMap{
        add(&minHeap, [2]int{key, val})
        
        if len(minHeap) > k {
            pop(&minHeap)
        }
    }
    
    for i := 0; i < len(minHeap); i++ {
        res = append(res, minHeap[i][0])
    }
    
    return res
}

func add(heap *[][2]int, item [2]int) {
    if heap == nil {
        panic("nil pointer")
    }    
     
    if len(*heap) == 0 {
        *heap = append(*heap, item)
        return
    }
    
    *heap = append(*heap, item)
    heapUp(*heap, len(*heap) - 1)
}

func pop(heap *[][2]int) [2]int {
    if heap == nil {
        panic("nil pointer")
    }    
    
    if len(*heap) == 0 {
        panic("empty heap")
    }
    
    poppedItem := (*heap)[0]
    
    lastIdx := len(*heap) - 1
    (*heap)[0], (*heap)[lastIdx] = (*heap)[lastIdx], (*heap)[0]
    *heap = (*heap)[:lastIdx]
    heapDown(*heap, 0, 0, lastIdx - 1)
    
    return poppedItem
}

func heapUp(nums [][2]int, pos int) {
    parent := (pos -  1)/2
    
    if parent >= 0 && nums[pos][1] < nums[parent][1] {
        nums[pos], nums[parent] = nums[parent], nums[pos]
        heapUp(nums, parent)
    }
}

func heapDown(nums [][2]int, pos int, low int, high int) {
    l := pos*2 + 1
    r := pos*2 + 2
    smaller := pos
    
    if low <= l && l <= high && nums[l][1] < nums[smaller][1] {
        smaller = l
    }
    
    if low <= r && r <= high && nums[r][1] < nums[smaller][1] {
        smaller = r
    }
    
    if smaller != pos {
        nums[smaller], nums[pos] = nums[pos], nums[smaller]
        heapDown(nums, smaller, 0, high)
    }
}