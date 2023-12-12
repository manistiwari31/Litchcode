func removeElements(head *ListNode, val int) *ListNode {
    if head == nil {
        return nil
    }

    if head.Val == val {
        return removeElements(head.Next, val)
    }

    var next *ListNode = head.Next
    if next != nil {
        if next.Val == val {
            next = removeElements(next.Next, val)
        } else {
            next.Next = removeElements(next.Next, val)
        }
    }
    head.Next = next

    return head
}