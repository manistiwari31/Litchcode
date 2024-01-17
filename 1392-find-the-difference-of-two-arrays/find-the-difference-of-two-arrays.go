func findDifference(nums1 []int, nums2 []int) [][]int {
    s1 := make(map[int]bool)
    s2 := make(map[int]bool)
    for _,num := range(nums1){
        s1[num] = true
    }
    for _,num := range(nums2){
        s2[num] = true
    }
    answer := make([][]int,2)
    for item,_ := range(s1){
        if _,ok := s2[item]; !ok{
            answer[0] = append(answer[0],item)
        }
    }
    for item,_ := range(s2){
        if _,ok := s1[item]; !ok{
            answer[1] = append(answer[1],item)
        }
    }
    return answer
}