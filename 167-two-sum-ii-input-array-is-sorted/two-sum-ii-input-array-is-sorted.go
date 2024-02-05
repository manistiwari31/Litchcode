func twoSum(numbers []int, target int) []int {
    start := 0
    end := len(numbers)-1

    for start < end {
        sum := numbers[start] + numbers[end]

        if sum == target {
            break
        }

        if sum > target {
            end--
        } else {
            start++
        }
    }
    return []int{start+1,end+1}
}
