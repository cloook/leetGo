/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	result := &ListNode{Next: head}
	lastSorted := head
	cur := head.Next
	for cur != nil {
		if cur.Val > lastSorted.Val {
			lastSorted = lastSorted.Next
		} else {
			prev := result
			for prev.Next.Val <= cur.Val {
				prev = prev.Next
			}
			lastSorted.Next = cur.Next
			cur.Next = prev.Next
			prev.Next = cur
		}
		cur = lastSorted.Next
	}
	return result.Next
}