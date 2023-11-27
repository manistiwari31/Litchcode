
type NestedIterator struct {
    flattenedList []int
    idx int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
    
    flattenedList := []int{}
    var flatten func(*[]int, []*NestedInteger)
    flatten = func(flattenedList *[]int, nestedList []*NestedInteger) {
        for _, ni := range nestedList {
            if ni.IsInteger() {
                 *flattenedList=append(*flattenedList,ni.GetInteger())
                continue
            }
            flatten(flattenedList, ni.GetList())
        }
    }
    flatten(&flattenedList, nestedList)
    return &NestedIterator{flattenedList:flattenedList}
}

func (this *NestedIterator) Next() int {
    if !this.HasNext() {
        return -1
    }
      val:=this.flattenedList[this.idx]
    this.idx++
    return val
}

func (this *NestedIterator) HasNext() bool {
    return this.idx<len(this.flattenedList)
}