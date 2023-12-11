func findSpecialInteger(arr []int) int {
    return  findAndCount(arr)
}

func findAndCount(num []int) int {
    max := 0
    res := 0
    n := len(num)
    for n>0{
        c:= 0
    for _, ele :=range(num){
        if num[n-1]==ele{
            c++
        }
    }
    if max<c{
        max = c
        res = num[n-1]
    }
    n--
    }
    return res
}