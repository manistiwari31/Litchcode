func findLeastNumOfUniqueInts(arr []int, k int) int {
    m := map[int]int{}

    for _, i := range arr {
        m[i]++
    }

    freq := make([]int, 0, len(m))
    for _, v := range m {
        freq = append(freq, v)
    }
    sort.Ints(freq)

    remaining := len(freq)
    for _, f := range freq {
        k -= f
        if k < 0 {
            break
        }
        remaining--
    }

    return remaining
}
