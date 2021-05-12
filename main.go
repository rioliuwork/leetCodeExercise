package main

import "leetCodeExercise/exercise"

func main() {
	var head, root *exercise.ListNode
	head = &exercise.ListNode{Val: 0}
	root = head
	for i := 1; i <= 3; i++ {
		root.Next = &exercise.ListNode{Val: i}
		root = root.Next
	}
	val := exercise.ReverseKGroup(head, 3)
	//print(val)
	for val != nil {
		print(val.Val)
		val = val.Next
	}
}
