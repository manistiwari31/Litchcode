func getIntersectionNode(headA, headB *ListNode) *ListNode {
    seen := make(map[*ListNode]bool)
    for n := headA; n != nil; n = n.Next {
        seen[n] = true
    }

    for n := headB; n != nil; n = n.Next {
        if seen[n] {
            return n
        }
    }
    return nil
}