package four

/*
给你链表的头结点head，请将其按 升序 排列并返回 排序后的链表 。

进阶：
	你可以在O(nlogn) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？

示例 1：
	输入：head = [4,2,1,3]
	输出：[1,2,3,4]

示例 2：
	输入：head = [-1,5,3,4,0]
	输出：[-1,0,3,4,5]

示例 3：
	输入：head = []
	输出：[]
*/
//func sortList(head *ListNode) *ListNode {
//	if head == nil {
//		return head
//	}
//	p := head
//	arr := []int{}
//	for p != nil {
//		arr = append(arr, p.Val)
//		p = p.Next
//	}
//	p = head
//	sort.Ints(arr)
//	for i, lenA := 0, len(arr); i < lenA; i++ {
//		p.Val = arr[i]
//		p = p.Next
//	}
//	return head
//}

// 不使用 api 方式
// 采用快排方式
func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	// 统计长度
	length := 0
	for node := head; node != nil; node = node.Next {
		length++
	}
	// 结果链表
	dummyHead := &ListNode{Next: head}


	return dummyHead.Next
}
func mergeNodeList(head1, head2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	temp, temp1, temp2 := dummyHead, head1, head2
	for temp1 != nil && temp2 != nil {
		if temp1.Val <= temp2.Val {
			temp.Next = temp1
			temp1 = temp1.Next
		} else {
			temp.Next = temp2
			temp2 = temp2.Next
		}
		temp = temp.Next
	}
	if temp1 != nil {
		temp.Next = temp1
	} else if temp2 != nil {
		temp.Next = temp2
	}
	return dummyHead.Next
}