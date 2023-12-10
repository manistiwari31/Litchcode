func transpose(matrix [][]int) [][]int {
   m:=len(matrix)
    n:=len(matrix[0])
    res:=make([][]int,n)
    for i:=0;i<n; i++{
        res[i]=make([]int,m)
    }
    for i, val := range(matrix){
           for j, v := range(val){
               res[j][i] = v
           }
        }
    return res       
}


     
