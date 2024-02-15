func largestPerimeter(nums []int) int64 {
    n:=len(nums)
    sort.Ints(nums)
    previousSum:=int64(0)
    ans:=int64(-1)
    for i:=0; i<n; i++ {
        if i>=2 && previousSum > int64(nums[i]){
            ans=int64(nums[i])+previousSum
        }
         previousSum+=int64(nums[i])
    }
    return ans
}