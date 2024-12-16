package config

import (
	"fmt"
)

// ListNode defines a node in the linked list
type ListNode struct {
	Value int
	Next  *ListNode
}

func CreateLinkedList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	head := &ListNode{Value: values[0]}
	current := head
	for _, val := range values[1:] {
		current.Next = &ListNode{Value: val}
		current = current.Next
	}
	return head
}

func PrintLinkedList(head *ListNode) {
	fmt.Print("\nlinked list: ")
	for head != nil {
		fmt.Printf("%d -> ", head.Value)
		head = head.Next
	}
	fmt.Println("nil")
}
