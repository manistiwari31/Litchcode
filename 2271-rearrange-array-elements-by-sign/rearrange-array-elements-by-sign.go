func rearrangeArray(nums []int) []int {
    var neg []int
    var pos []int

    for _, i := range nums {
        if i>0 {
            neg = append(neg, i)
        } else {
            pos = append(pos, i)
        }
    }
    var res []int
    for i := range pos {
        res = append(res, neg[i])
        res = append(res,pos[i])
     
    }


    return res
}