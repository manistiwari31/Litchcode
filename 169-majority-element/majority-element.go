func majorityElement(nums []int) int {
    key := make(map[int]int)
    for _,n := range nums {key[n]++}

    res := 0
    max := 0

    for i,n := range key {
        if n>max {
            max = n
            res = i
        }
    }
    return res
}