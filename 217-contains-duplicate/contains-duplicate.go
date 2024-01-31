func containsDuplicate(nums []int) bool {
    hashset := make(map[int]bool)

    for _,n := range nums {
        fmt.Print(n)
        if hashset[n]{
            return true
        }

        hashset[n] = true
    }
    return false
}