func countNicePairs(nums []int) int {
       if(len(nums)<2){
        return 0
    }
    dictReverse:=make(map[int]int)
    countPairs:=0 
    moduloNumber := 1_000_000_007
    for i:=0;i<len(nums);i++{
        diff:= nums[i] - reverseNum(nums[i])
        if existingPairs,ok:=dictReverse[diff];ok{
            countPairs+=existingPairs
        }
        dictReverse[diff]+=1
    }
    return countPairs%moduloNumber
}


func reverseNum(n int) int {
    new_int := 0
    for n > 0 {
        remainder := n % 10
        new_int *= 10
        new_int += remainder 
        n /= 10
    }
    return new_int 
}