package exercise

/**
给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。

k 是一个正整数，它的值小于或等于链表的长度。

如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

进阶：

你可以设计一个只使用常数额外空间的算法来解决此问题吗？
你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。


示例 1：


输入：head = [1,2,3,4,5], k = 2
输出：[2,1,4,3,5]
示例 2：


输入：head = [1,2,3,4,5], k = 3
输出：[3,2,1,4,5]
示例 3：

输入：head = [1,2,3,4,5], k = 1
输出：[1,2,3,4,5]
示例 4：

输入：head = [1], k = 1
输出：[1]
提示：

列表中节点的数量在范围 sz 内
1 <= sz <= 5000
0 <= Node.val <= 1000
1 <= k <= sz
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func ReverseKGroup(head *ListNode, k int) *ListNode {
	var child, root, pointer, nextPointer, endPointer, flagPointer, test *ListNode
	nextPointer = head
	root = head
	flag := 1
	use := true
	for nextPointer != nil {
		child = nil
		test = nextPointer
		for j := 1; j < k; j++ {
			test = test.Next
			if test == nil {
				child = nextPointer
				nextPointer = nil
				use = false
				break
			}
		}
		if use {
			for i := 0; i < k; i++ {
				pointer = nextPointer
				nextPointer = nextPointer.Next
				pointer.Next = child
				child = pointer
				if i == 0 {
					endPointer = child
				}
				if nextPointer == nil {
					break
				}
			}
		}
		if flag == 1 {
			root = child
			flag++
		} else {
			flagPointer.Next = child
		}
		flagPointer = endPointer
	}
	flagPointer = nil
	return root
}
