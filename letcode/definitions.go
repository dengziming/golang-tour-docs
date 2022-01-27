package letcode

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func createListNode(nodes []int) *ListNode {
	if len(nodes) == 1 {
		return &ListNode{nodes[0], nil}
	}
	return &ListNode{nodes[0], createListNode(nodes[1:])}
}

func printlnListNode(node *ListNode) {
	if node == nil {
		println()
		return
	}
	fmt.Print(node.Val)
	fmt.Print(" ")
}

func printlnArray(data [][]int){
	if len(data) == 0 {
		println("nil")
	}
	for _, ele := range data {
		for _, d := range ele {
			print(d)
			print(" ")
		}
		println()
	}
}
