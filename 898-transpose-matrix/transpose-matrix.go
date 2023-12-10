func transpose(matrix [][]int) [][]int {
    r, c := len(matrix), len(matrix[0])
    
    res := make([][]int, c)
    for rid := 0; rid < c; rid++ {
        res[rid] = make([]int, r)
    }

    for rid := 0; rid < r; rid++ {
        for cid := 0; cid < c; cid++ {
            res[cid][rid] = matrix[rid][cid]
        }
    }
    
    return res
}