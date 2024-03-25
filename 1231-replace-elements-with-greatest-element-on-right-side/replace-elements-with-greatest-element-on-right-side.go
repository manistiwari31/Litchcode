func replaceElements(arr []int) []int {
    for i:= range arr {
        arr[i]= findMax(arr[i+1:])
    }
    return arr
}

func findMax(arr[]int) int {
    maX := -1
    if len(arr)==0 {
        return maX
    }
    for i:= range arr {
        if arr[i]>maX {
            maX = arr[i]
        }
    }
    return maX
}