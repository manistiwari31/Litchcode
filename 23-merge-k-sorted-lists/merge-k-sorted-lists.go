/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 { return  nil}

    for len(lists) > 1 {
        mergedList := []*ListNode{}

        for i := 0; i < len(lists); i += 2 {
            if i+1 >= len(lists) {
                mergedList = append(mergedList, lists[i]) 
                continue
            }

            mergedList = append(mergedList, mergeTwoList(lists[i], lists[i+1]))
        }

        lists = mergedList
    }

    return lists[0]
}

func mergeTwoList(list1 *ListNode, list2 *ListNode) *ListNode {
    if list1 == nil { return list2 }
    if list2 == nil { return list1 }

    res := new(ListNode)
    tail := res

    for list1 != nil && list2 != nil {
        if list1.Val <= list2.Val {
            tail.Next = list1
            list1 = list1.Next
        } else {
            tail.Next = list2
            list2 = list2.Next
        }

        tail = tail.Next
    }

    if list1 != nil { tail.Next = list1 }
    if list2 != nil { tail.Next = list2 }

    return res.Next
}