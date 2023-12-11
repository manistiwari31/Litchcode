func findSpecialInteger(arr []int) int {
    
    max := 0
    res := 0
    n := len(arr)
    for n>0{
        c:= 0
    for _, ele :=range(arr){
        if arr[n-1]==ele{
            c++
        }
    }
    if max<c{
        max = c
        res = arr[n-1]
    }
    n--
    }
    return res
}