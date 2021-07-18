package reverse_nodes_in_k_group

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	t1 := head
	i := 0
	var prev, nn *ListNode
	curr := head
	for (i < k) && (t1 != nil) {
		t1 = t1.Next
		i++
	}
	if i != k {
		return head
	}
	i = 0
	for i < k {
		nn = curr.Next
		curr.Next = prev
		prev = curr
		curr = nn
		i++
	}
	if nn != nil {
		head.Next = reverseKGroup(curr, k)
	}

	return prev

}
